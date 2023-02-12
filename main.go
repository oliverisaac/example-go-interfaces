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

	// You cannot pass things to Results() because it expects a slice of interfaces
	// but it is a slice of structs:
	// There is a workaround using generics:
	// https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices
	fmt.Println(firstProcessor.Results(things))
	fmt.Println(secondProcessor.Results(things))
}
