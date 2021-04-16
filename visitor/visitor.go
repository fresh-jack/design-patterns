package visitor

import "fmt"

type IVisitor interface {
	Visit()
}

type ProductionVisitor struct {
}

func (p ProductionVisitor) Visit() {
	fmt.Println("production")
}

type TestingVisitor struct {
}

func (t TestingVisitor) Visit() {
	fmt.Println("testing")
}

type IElement interface {
	Accept(IVisitor)
}

type Element struct {
}

func (e Element) Accept(v IVisitor) {
	v.Visit()
}

type EnvExample struct {
	Element
}

func (e EnvExample) Print(v IVisitor) {
	e.Element.Accept(v)
}

func Test() {
	e := Element{}
	e.Accept(new(ProductionVisitor))
	e.Accept(new(TestingVisitor))
}
