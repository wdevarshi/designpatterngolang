package encapsulation

import "fmt"

type Fooer interface {
	Foo1()
	Foo2()
	Foo3()
}

type Foo struct {
}

func (f Foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f Foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f Foo) Foo3() {
	fmt.Println("Foo3() here")
}