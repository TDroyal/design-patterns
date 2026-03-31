package decoratorpattern

import (
	"context"
	"fmt"
)

type Food interface {
	Eat() string
	Cost() float32
}

// 主类
type Rice struct{}

func NewRice() Food {
	return &Rice{}
}

func (r *Rice) Eat() string {
	return "吃香喷喷的大米饭"
}

func (r *Rice) Cost() float32 {
	return 1
}

// 装饰类
type Decorator Food

// laoganma 每个装饰器类的作用是对食物进行一轮装饰增强，因此需要在构造器函数中传入待装饰的食物，
// 然后通过重写食物的 Eat 和 Cost 方法（在重写方法中增加额外的逻辑），实现对应的增强装饰效果.
type LaoGanMaDecorator struct {
	Decorator
}

func NewLaoGanMaDecorator(d Decorator) Decorator { // 返回加入老干妈装饰后的食物
	return &LaoGanMaDecorator{
		Decorator: d,
	}
}

func (l *LaoGanMaDecorator) Eat() string {
	return "加一份老干妈\n" + l.Decorator.Eat() // 执行主类方法前先执行装饰器的逻辑
}

func (l *LaoGanMaDecorator) Cost() float32 {
	return 1 + l.Decorator.Cost()
}

// fried egg
type FriedEggDecorator struct {
	Decorator
}

func NewFriedEggDecorator(d Decorator) Decorator {
	return &FriedEggDecorator{
		Decorator: d,
	}
}

func (l *FriedEggDecorator) Eat() string {
	return "加一个煎蛋\n" + l.Decorator.Eat()
}

func (l *FriedEggDecorator) Cost() float32 {
	return 1.5 + l.Decorator.Cost()
}

// 闭包实现装饰增强函数的实现示例，其实现也是遵循着装饰器模式的思路，但在形式上会更加简洁直观一些

type handleFunc func(ctx context.Context, param map[string]interface{}) error

func Decorate(fn handleFunc) handleFunc {
	return func(ctx context.Context, param map[string]interface{}) error {
		// 前处理
		fmt.Println("preprocess...")
		err := fn(ctx, param)
		// 后处理
		fmt.Println("postprocess...")
		return err
	}
}
