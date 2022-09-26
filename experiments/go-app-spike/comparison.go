package main

import "fmt"

type foo struct {
	Foo       string
	evaluated string
}

func (f *foo) evaluate() *foo {
	f.evaluated = fmt.Sprintf("I have set the prop: %v and am evaluated", f.Foo)
	return f

}
