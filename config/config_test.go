package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  error
	}{
		{name: "first", input: "https://www.nike.com", want: nil},
		{name: "second ", input: "https://geekbrains.com", want: nil},
		{name: "third", input: "vk", want: nil},
	}

	for _, tc := range tests {
		got := urlValidation(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}

func TestSomething(t *testing.T) {
	assert.Equal(t, nil, urlValidation("https://www.nike.com"), "they should be equal")
}
