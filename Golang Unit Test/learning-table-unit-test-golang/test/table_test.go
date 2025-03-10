package test

import (
	"learning-table-unit-test-golang/helper"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTableSayHello(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Abdan",
			request:  "Abdan",
			expected: "Hello Abdan",
		},
		{
			name:     "Zaki",
			request:  "Zaki",
			expected: "Hello Zaki",
		},
		{
			name:     "Alifian",
			request:  "Alifian",
			expected: "Hello Alifian",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sayHello := helper.SayHello(test.request)
			require.Equal(t, test.expected, sayHello)
		})
	}
}
