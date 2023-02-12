package firstProcessor

import "strings"

type Processable interface {
	// Value will be used to generate the output
	Value() string
	// Uppercase returns true when the value should be uppercased
	Uppercase() bool
}

// We use generics here to workaround "array of interfaces" limitations:
// https://dusted.codes/using-go-generics-to-pass-struct-slices-for-interface-slices
func ResultsGenerics[P Processable](items []P) string {
	result := []string{}
	for _, item := range items {
		v := item.Value()
		if item.Uppercase() {
			v = strings.ToUpper(v)
		}
		result = append(result, v+"\n")
	}
	return strings.Join(result, "")
}

func ResultsInterfaces(items []Processable) string {
	return ResultsGenerics(items)
}
