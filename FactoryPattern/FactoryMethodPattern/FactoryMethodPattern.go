package factorymethodpattern

import "fmt"

type Fruit interface {
	Eat()
}

// strawberry
type Strawberry struct {
	name string
}

func (s *Strawberry) Eat() {
	fmt.Printf("I am strawberry: %s\n", s.name)
}

func NewStrawberry(name string) Fruit {
	return &Strawberry{
		name: name,
	}
}

// Cherry
type Cherry struct {
	name string
}

func (s *Cherry) Eat() {
	fmt.Printf("I am Cherry: %s\n", s.name)
}

func NewCherry(name string) Fruit {
	return &Cherry{
		name: name,
	}
}

// fruit factory
type FruitFactory interface {
	CreateFruit(string) Fruit
}

// Strawberry factory
type StrawberryFactory struct{}

func (s *StrawberryFactory) CreateFruit(name string) Fruit {
	return NewStrawberry(name)
}

func NewStrawberryFactory() FruitFactory {
	return &StrawberryFactory{}
}

// Cherry factory
type CherryFactory struct{}

func (s *CherryFactory) CreateFruit(name string) Fruit {
	return NewCherry(name)
}

func NewCherryFactory() FruitFactory {
	return &CherryFactory{}
}

/*
	解决了扩展水果类不满足开闭原则的问题
*/

/*
	p1.需要为每个水果单独实现一个工厂类,代码冗余度较高
	p2.原本构造多个水果类时存在的公共切面不复存在，一些通用的逻辑需要在每个水果工厂实现类中重复声明一遍
*/
