package main

import (
	"./llrb"
	"fmt"
)

type Person struct {
	age  int
	name string
}

func CmpAge(a, b interface{}) bool {
	return a.(Person).age < b.(Person).age
}

func main() {
	joe := Person{23, "Joe"}
	ann := Person{45, "Ann"}
	jim := Person{37, "Jim"}
	jac := Person{43, "Jac"}
	tom := Person{84, "Tom"}
	tim := Person{12, "Tim"}

	bob := Person{90, "Bob"}

	tree := llrb.New(CmpAge)

	tree.Insert(joe)
	tree.Insert(ann)
	tree.Insert(jim)
	tree.Insert(jac)
	tree.Insert(tom)
	tree.Insert(tim)

	fmt.Printf("%v\n", tree.Search(joe))
	fmt.Printf("%v\n", tree.Search(ann))
	fmt.Printf("%v\n", tree.Search(bob))
	fmt.Printf("%v\n", tree.Max())
	fmt.Printf("%v\n", tree.Min())
}
