package main

import "github.com/eduwr/go-algorithms/lib/linkedstack"

func main() {
	l := linkedstack.LinkedStackOfStrings{}

	println(l.IsEmpty())

	l.Push("Eduardo")
	l.Push("Wronscki")
	l.Push("Ricardo")
	println(l.Pop())
	println(l.Pop())
	println(l.Pop())
}
