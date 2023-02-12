package main

import (
	firstProcessor "github.com/oliverisaac/example-go-interfaces/first-processor"
	secondProcessor "github.com/oliverisaac/example-go-interfaces/second-processor"
)

type MyThing struct {
	value string
}

// We use this to ensure that MyThing implements the interfaces
var _ firstProcessor.Processable = MyThing{}
var _ secondProcessor.Processable = MyThing{}

func (m MyThing) Value() string {
	return m.value
}

func (m MyThing) Quantity() int {
	return len(m.value)
}

func (m MyThing) Uppercase() bool {
	return len(m.value)%2 == 0
}
