package simplefactorypattern

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
type FruitFactory struct{}

func NewFruitFactory() *FruitFactory {
	return &FruitFactory{}
}

func (f *FruitFactory) CreateFruit(typ string, name string) (Fruit, error) {
	switch typ {
	case "Strawberry":
		return NewStrawberry(name), nil
	case "Cherry":
		return NewCherry(name), nil
	default:
		return nil, fmt.Errorf("fruit type: %s is not supported", typ)
	}
}

/*
	p1.不符合开闭原则（面向扩展开放，面向修改关闭）
	p2.当需要支持的水果类型 typ 数量提升时，这个 CreateFruit 方法会存在方法圈复杂度过高的问题
*/

// solve p2
type fruitConstructor func(name string) Fruit

type FruitFactoryV2 struct {
	c map[string]fruitConstructor
}

func NewFruitFactoryV2() *FruitFactoryV2 {
	return &FruitFactoryV2{
		c: map[string]fruitConstructor{
			"Strawberry": NewStrawberry,
			"Cherry":     NewCherry,
		},
	}
}

func (f *FruitFactoryV2) CreateFruitV2(typ, name string) (Fruit, error) {
	fruitconstructor, ok := f.c[typ]
	if !ok {
		return nil, fmt.Errorf("fruit type: %s is not supported", typ)
	}
	return fruitconstructor(name), nil
}
