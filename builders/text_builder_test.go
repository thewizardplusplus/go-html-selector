package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextBuilder_TextParts(test *testing.T) {
	builder := TextBuilder{
		textParts: [][]byte{
			[]byte("text part #1"),
			[]byte("text part #2"),
		},
	}
	gotTextParts := builder.TextParts()

	assert.Equal(test, builder.textParts, gotTextParts)
}
