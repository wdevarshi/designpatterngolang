package encapsulation

import "fmt"

type foo struct {
}

func (f foo) Foo1() {
	fmt.Println("Foo1() here")
}

func (f foo) Foo2() {
	fmt.Println("Foo2() here")
}

func (f foo) Foo3() {
	fmt.Println("Foo3() here")
}

func NewFoo() Fooer {
	return &Foo{}
}

/*to hide the Foo type above and expose just the interface you could rename it to lower case foo and
provide a NewFoo() function that returns the public Fooer interface:*/

//f := NewFoo()
//f.Foo1()