package decoratorpattern

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	rice := NewRice()
	fmt.Println(rice.Eat(), rice.Cost())
	fmt.Println()

	// 加一份老干妈
	rice = NewLaoGanMaDecorator(rice)
	fmt.Println(rice.Eat(), rice.Cost())
	fmt.Println()

	// 再加一个煎蛋
	rice = NewFriedEggDecorator(rice)
	fmt.Println(rice.Eat(), rice.Cost())
	fmt.Println()

	r := NewFriedEggDecorator(NewLaoGanMaDecorator(NewFriedEggDecorator(NewRice())))
	fmt.Println(r.Eat(), r.Cost())
}

/*
=== RUN   TestDecorator
吃香喷喷的大米饭 1

加一份老干妈
吃香喷喷的大米饭 2

加一个煎蛋
加一份老干妈
吃香喷喷的大米饭 3.5

加一个煎蛋
加一份老干妈
加一个煎蛋
吃香喷喷的大米饭 5
--- PASS: TestDecorator (0.00s)
PASS
*/
