package morestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseRunes(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseRunes(c.in)
		assert.Equal(got, c.want)
	}
}
