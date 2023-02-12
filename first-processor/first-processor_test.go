package firstProcessor

import "testing"

type mockProcessable struct {
	value   string
	doUpper bool
}

func (m mockProcessable) Value() string {
	return m.value
}

func (m mockProcessable) Uppercase() bool {
	return m.doUpper
}

func newMockProcessable(value string, doUpper bool) Processable {
	return mockProcessable{
		value:   value,
		doUpper: doUpper,
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
				newMockProcessable("one", false),
			},
			want: "one\n",
		},
		{
			name: "multiple items",
			items: []Processable{
				newMockProcessable("one", false),
				newMockProcessable("two", true),
			},
			want: "one\nTWO\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Results(tt.items); got != tt.want {
				t.Errorf("Results() = %v, want %v", got, tt.want)
			}
		})
	}
}
