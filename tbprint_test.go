package tbprint

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxWidths(t *testing.T) {
	assert := assert.New(t)
	var testdata = []struct {
		input  [][]string
		result map[int]int
	}{
		{
			[][]string{
				[]string{"a", "bcd", "ef"},
			},
			map[int]int{
				0: 1,
				1: 3,
				2: 2,
			},
		},
		{
			[][]string{
				[]string{"a", "bcd", "ef"},
				[]string{"abc", "defg"},
			},
			map[int]int{
				0: 3,
				1: 4,
				2: 2,
			},
		},
		{
			[][]string{
				[]string{"a", "bcd", "ef"},
				[]string{"abc", "defg"},
				[]string{"a", "bc", "def", "fg", "i"},
			},
			map[int]int{
				0: 3,
				1: 4,
				2: 3,
				3: 2,
				4: 1,
			},
		},
	}

	for _, test := range testdata {
		widths := maxWidths(test.input)
		t.Log("widths =", widths)
		for idx, value := range widths {
			assert.Equal(test.result[idx], value)
		}
	}
}

func TestFprint(t *testing.T) {
	var testdata = []struct {
		input [][]string
	}{
		{
			[][]string{
				[]string{"", "Name", "Age", "City"},
				[]string{"", "Yunhong", "23", "Seoul"},
				[]string{"", "Hyunjung", "23", "New York"},
				[]string{"", "Joonhyun", "1", "New York"},
				[]string{"Total", "--", "48", "--"},
			},
		},
	}

	for _, test := range testdata {
		var b = bytes.Buffer{}
		writer := bufio.NewWriter(&b)
		Fprint(writer, test.input)
		writer.Flush()
		t.Log(b.String())
	}
}
