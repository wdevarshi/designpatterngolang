package main
/*
  A
 / \
B   C
 \ /
  D
*/

type B struct {

}

type C struct {

}

func (b B) Foo() {}

func (c C) Foo() {}

type D struct {
	B
	C
}

func main() {
	d := D{B{}, C{}}
	// d.Foo() // <- ambiguous
	d.B.Foo()  // <- ok
	d.C.Foo()  // <- ok
}

//For further reading; Please refer virtual inheritance