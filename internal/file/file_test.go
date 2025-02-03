package file

import (
	"slices"
	"testing"
)

func TestDeleteEmptyLines(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{
			[]string{"foo", "bar", "buz"},
			[]string{"foo", "bar", "buz"},
		},
		{
			[]string{"\tfoo", " bar", ""},
			[]string{"\tfoo", " bar"},
		},
		{
			[]string{"", "", "foo", "", "bar"},
			[]string{"foo", "bar"},
		},
		{
			[]string{"\t", "    ", "\n"},
			[]string{},
		},
		{
			[]string{"\t\n   ", "foo"},
			[]string{"foo"},
		},
	}

	for i, test := range tests {
		result := deleteEmptyLines(test.input)
		if slices.Compare(result, test.expected) != 0 {
			t.Fatalf("Test %d failed:\n\tgot=%v\n\texpected=%v", i, result, test.expected)
		}
	}
}
