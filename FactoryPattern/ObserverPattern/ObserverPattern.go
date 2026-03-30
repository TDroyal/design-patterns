package observerpattern

import (
	"context"
	"fmt"
	"sync"
)

type Event struct {
	Topic string
	Val   any // 变更详情
}

type Observer interface {
	Onchange(ctx context.Context, e *Event) error // 当订阅的topic有事件发生时，调用对应的observer的Onchange
}

type EventBus interface {
	Subscribe(topic string, o Observer)    // 注册observer
	Unsubscribe(topic string, o Observer)  // 注销observer
	Publish(ctx context.Context, e *Event) // 发布event，提醒对应topic的observers触发Onchange
}

// demo observer
type BaseObserver struct {
	name string
}

func NewBaseObserver(name string) *BaseObserver {
	return &BaseObserver{
		name: name,
	}
}

func (b *BaseObserver) Onchange(ctx context.Context, e *Event) error {
	// ...
	fmt.Printf("observer: %s, event topic: %s, event val: %v\n", b.name, e.Topic, e.Val)

	return fmt.Errorf("observer: %s, event topic: %s, event val: %v", b.name, e.Topic, e.Val) // return nil
}

// demo eventbus
type void struct{}

type BaseEventBus struct {
	observers map[string]map[Observer]void // 维护好observers关注的topic    map[Observer]void 类比一个set
	mu        sync.RWMutex
}

func NewBaseEventBus() BaseEventBus {
	return BaseEventBus{
		observers: make(map[string]map[Observer]void),
		mu:        sync.RWMutex{},
	}
}

func (b *BaseEventBus) Subscribe(topic string, o Observer) {
	b.mu.Lock()
	if _, exist := b.observers[topic]; !exist {
		b.observers[topic] = make(map[Observer]void)
	}
	b.observers[topic][o] = void{}
	b.mu.Unlock()
}

func (b *BaseEventBus) Unsubscribe(topic string, o Observer) {
	b.mu.Lock()
	if obs, exist := b.observers[topic]; exist {
		if _, existAgain := obs[o]; existAgain {
			delete(b.observers[topic], o)
		}
	}
	b.mu.Unlock()
}

// async events publish demo

type observerWithErr struct {
	o   Observer
	err error
}

type AsyncEventBus struct {
	BaseEventBus
	errChan chan *observerWithErr // error channel (async handle)
	ctx     context.Context
	stop    context.CancelFunc
}

func NewAsyncEventBus() *AsyncEventBus {
	aeBus := AsyncEventBus{
		BaseEventBus: NewBaseEventBus(),
	}
	aeBus.ctx, aeBus.stop = context.WithCancel(context.Background())
	aeBus.errChan = make(chan *observerWithErr, 10) // ?
	go aeBus.handleErr()
	return &aeBus
}

func (a *AsyncEventBus) Publish(ctx context.Context, e *Event) {
	a.mu.RLock()
	observers, exist := a.observers[e.Topic]
	a.mu.RUnlock()

	if !exist {
		return
	}

	for o := range observers {
		go func() {
			if err := o.Onchange(ctx, e); err != nil {
				select {
				case <-a.ctx.Done():
					return
				case a.errChan <- &observerWithErr{o, err}:
				}
			}
		}()
	}
}

func (a *AsyncEventBus) Stop() {
	a.stop()
}

func (a *AsyncEventBus) handleErr() {
	for {
		select {
		case <-a.ctx.Done():
			return
		case observerErr := <-a.errChan:
			// handle err
			fmt.Printf("handle err -> observer: %v, err: %v\n", observerErr.o, observerErr.err)
		}
	}
}
