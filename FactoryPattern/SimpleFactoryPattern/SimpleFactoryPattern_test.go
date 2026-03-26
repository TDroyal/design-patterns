package simplefactorypattern

import (
	"testing"
)

func TestSimpleFactoryPattern(t *testing.T) {
	fruitFactory := NewFruitFactoryV2()
	strawberry, _ := fruitFactory.CreateFruitV2("Strawberry", "1")
	strawberry.Eat()

	cherry, _ := fruitFactory.CreateFruitV2("Cherry", "2")
	cherry.Eat()

	orange, err := fruitFactory.CreateFruitV2("Orange", "3")
	if err != nil {
		t.Error(err)
		return
	}
	orange.Eat()
}

/*
=== RUN   TestSimpleFactoryPattern
I am strawberry: 1
I am Cherry: 2
    e:\project\DesignPatterns\FactoryPattern\SimpleFactoryPattern\SimpleFactoryPattern_test.go:17: fruit type: Orange is not supported
--- FAIL: TestSimpleFactoryPattern (0.00s)
FAIL
FAIL    DesignPatterns/FactoryPattern/SimpleFactoryPattern      0.114s
*/
