package main

import (
	"fmt"

	as "github.com/chyld/goxercise/arrayslice"
	"github.com/chyld/goxercise/control"
	"github.com/chyld/goxercise/primitives"
)

func main() {
	fmt.Println("{goxercise}")
	fmt.Println("Sum:", primitives.Sum(3, 4))
	fmt.Println("Double:", primitives.Double(3))
	fmt.Println("Concat:", primitives.Concat("a", `b`))
	fmt.Println("AgeCheck:", control.AgeCheck(20))
	fmt.Println("Language:", control.Language(10))
	fmt.Println("Looping:", control.Looping(3, "go"))
	fmt.Println("Switching:", control.Switching(`android`))
	a, b := as.BasicArray()
	fmt.Printf("BasicArray: %#v %#v\n", a, b)
	as.BasicSlice()
}
