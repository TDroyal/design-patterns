package observerpattern

import (
	"context"
	"testing"
	"time"
)

func TestObserver(t *testing.T) {

	o1 := NewBaseObserver("royal_111")
	o2 := NewBaseObserver("hyl")
	o3 := NewBaseObserver("panghu")
	o4 := NewBaseObserver("laopo")
	o5 := NewBaseObserver("baobao")

	aBus := NewAsyncEventBus()
	defer aBus.Stop()

	aBus.Subscribe("couple", o1)
	aBus.Subscribe("couple", o2)
	aBus.Subscribe("others", o3)
	aBus.Subscribe("others", o4)
	aBus.Subscribe("pipi", o5)

	e1 := &Event{
		Topic: "couple",
		Val:   "you are a couple",
	}

	e2 := &Event{
		Topic: "others",
		Val:   "nothing",
	}
	aBus.Publish(context.Background(), e1)
	aBus.Publish(context.Background(), e2)

	<-time.After(time.Second * 3)
}

/*
=== RUN   TestObserver
observer: panghu, event topic: others, event val: nothing
observer: hyl, event topic: couple, event val: you are a couple
handle err -> observer: &{panghu}, err: observer: panghu, event topic: others, event val: nothing
handle err -> observer: &{hyl}, err: observer: hyl, event topic: couple, event val: you are a couple
observer: royal_111, event topic: couple, event val: you are a couple
observer: laopo, event topic: others, event val: nothing
handle err -> observer: &{royal_111}, err: observer: royal_111, event topic: couple, event val: you are a couple
handle err -> observer: &{laopo}, err: observer: laopo, event topic: others, event val: nothing
--- PASS: TestObserver (3.00s)
PASS
*/
