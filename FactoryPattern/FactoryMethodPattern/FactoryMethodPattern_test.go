package factorymethodpattern

import "testing"

func TestFactoryMethodPattern(t *testing.T) {
	strawberryFactory := NewStrawberryFactory()
	strawberry := strawberryFactory.CreateFruit("1")
	strawberry.Eat()

	cherryFactory := NewCherryFactory()
	cherry := cherryFactory.CreateFruit("2")
	cherry.Eat()
}

/*
=== RUN   TestFactoryMethodPattern
I am strawberry: 1
I am Cherry: 2
--- PASS: TestFactoryMethodPattern (0.00s)
PASS
ok      DesignPatterns/FactoryPattern/FactoryMethodPattern      1.267s
*/
