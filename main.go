package main

import (
	"fmt"

	firstProcessor "github.com/oliverisaac/example-go-interfaces/first-processor"
	secondProcessor "github.com/oliverisaac/example-go-interfaces/second-processor"
)

func main() {
	things := []MyThing{
		MyThing{value: "one"},
		MyThing{value: "two"},
		MyThing{value: "three"},
		MyThing{value: "four"},
	}

	doGenerics(things)
	doInterfaces(things)
	// This function shows something that doesn't work:
	// doCommonInterfaces(things)
}

func doGenerics(things []MyThing) {
	// You cannot pass things to Results() because it expects a slice of interfaces
	// but it is a slice of structs:
	// There is a workaround using generics:
	// https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices
	fmt.Println("Using generics:")
	fmt.Println("First processor:")
	fmt.Println(firstProcessor.ResultsGenerics(things))
	fmt.Println("Second processor:")
	fmt.Println(secondProcessor.ResultsGenerics(things))
}

func doInterfaces(things []MyThing) {
	// Alternatively, you can use the interface version by converting the array of structs to an array of interfaces:
	// This is an O(n) operation which is why Go doesn't do it automatically
	// and would be annoying if you had 10,000+ items
	fmt.Println("Using interfaces:")
	firstProcessables := make([]firstProcessor.Processable, len(things))
	for i, thing := range things {
		firstProcessables[i] = thing
	}
	// This is a compiler error:
	// firstProcessor.ResultsInterfaces(things)
	fmt.Println("First processor:")
	firstProcessor.ResultsInterfaces(firstProcessables)

	secondProcessables := make([]secondProcessor.Processable, len(things))
	for i, thing := range things {
		secondProcessables[i] = thing
	}
	// This is a compiler error:
	// secondProcessor.ResultsInterfaces(things)
	fmt.Println("S processor:")
	secondProcessor.ResultsInterfaces(secondProcessables)
}

/*
// This code does not work as Go does not let you convert a common "parent" interface into
// a child interface. This is a compiler error.
// https://www.reddit.com/r/golang/comments/crzqdp/comment/exd3wtk/?utm_source=share&utm_medium=web2x&context=3
func doCommonInterfaces(things []MyThing) {
	// Make a common interface that both processors use
	type CommonProcessable interface {
		firstProcessor.Processable
		secondProcessor.Processable
	}

	// Do the O(n) conversion once
	processsables := make([]CommonProcessable, len(things))
	for i, thing := range things {
		processsables[i] = thing
	}

	// Call our functions
	firstProcessor.ResultsInterfaces(processsables)
	secondProcessor.ResultsInterfaces(processsables)
}

*/
