package main

import "fmt"

type Product interface {
	Operation() string
}

type ConcreteProductA struct{} //具体产品类1

func (p *ConcreteProductA) Operation() string {
	return "ConcreteProductA operation"
}

type ConcreteProductB struct{} //具体产品类2

func (p *ConcreteProductB) Operation() string {
	return "ConcreteProductB operation"
}

type SimpleFactory struct{} //简单工厂类

func (f *SimpleFactory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func main() {
	factory := &SimpleFactory{}
	productA := factory.CreateProduct("A")
	fmt.Println(productA.Operation())
	fmt.Printf("%v", productA)

	productB := factory.CreateProduct("B")
	fmt.Println(productB.Operation())
	fmt.Printf("%v", productB)
}
