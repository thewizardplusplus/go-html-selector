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

func TestTextBuilder_AddText(test *testing.T) {
	builder := TextBuilder{
		textParts: [][]byte{
			[]byte("text part #1"),
			[]byte("text part #2"),
		},
	}
	builder.AddText([]byte("text part #3"))

	wantTextParts := [][]byte{
		[]byte("text part #1"),
		[]byte("text part #2"),
		[]byte("text part #3"),
	}
	assert.Equal(test, wantTextParts, builder.textParts)
}
