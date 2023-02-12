package secondProcessor

import "strings"

type Processable interface {
	// Value will be used to generate the output
	Value() string
	// Quantity returns the number of times the value should be repeated
	Quantity() int
}

// We use generics here to workaround "array of interfaces" limitations:
// https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices
func ResultsGenerics[P Processable](items []P) string {
	result := []string{}
	for _, item := range items {
		thisLine := []string{}
		for i := 0; i < item.Quantity(); i++ {
			thisLine = append(thisLine, item.Value())
		}
		result = append(result, strings.Join(thisLine, " ")+"\n")
	}
	return strings.Join(result, "")
}

func ResultsInterfaces(items []Processable) string {
	return ResultsGenerics(items)
}
