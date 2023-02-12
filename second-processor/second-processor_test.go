package secondProcessor

import "testing"

type mockProcessable struct {
	value    string
	quantity int
}

func (m mockProcessable) Value() string {
	return m.value
}

func (m mockProcessable) Quantity() int {
	return m.quantity
}

func newMockProcessable(value string, quantity int) mockProcessable {
	return mockProcessable{
		value:    value,
		quantity: quantity,
	}
}

func TestResults(t *testing.T) {
	tests := []struct {
		name  string
		items []Processable
		want  string
	}{
		{
			name:  "no items",
			items: []Processable{},
			want:  "",
		},
		{
			name: "one item",
			items: []Processable{
				newMockProcessable("one", 1),
			},
			want: "one\n",
		},
		{
			name: "multiple items",
			items: []Processable{
				newMockProcessable("one", 1),
				newMockProcessable("two", 2),
			},
			want: "one\ntwo two\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ResultsGenerics(tt.items); got != tt.want {
				t.Errorf("Results() = %v, want %v", got, tt.want)
			}
		})
	}
}
