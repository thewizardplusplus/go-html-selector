package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenBuilder_Attributes(test *testing.T) {
	builder := FlattenBuilder{
		attributes: [][]byte{
			[]byte("http://example.com/1"),
			[]byte("http://example.com/2"),
		},
	}
	gotAttributes := builder.Attributes()

	assert.Equal(test, builder.attributes, gotAttributes)
}

func TestFlattenBuilder_AddTag(test *testing.T) {
	builder := FlattenBuilder{
		attributes: [][]byte{
			[]byte("http://example.com/1"),
			[]byte("http://example.com/2"),
		},
	}
	builder.AddTag([]byte("a"))

	wantAttributes := [][]byte{
		[]byte("http://example.com/1"),
		[]byte("http://example.com/2"),
	}
	assert.Equal(test, wantAttributes, builder.attributes)
}

func TestFlattenBuilder_AddAttribute(test *testing.T) {
	builder := FlattenBuilder{
		attributes: [][]byte{
			[]byte("http://example.com/1"),
			[]byte("http://example.com/2"),
		},
	}
	builder.AddAttribute([]byte("href"), []byte("http://example.com/3"))

	wantAttributes := [][]byte{
		[]byte("http://example.com/1"),
		[]byte("http://example.com/2"),
		[]byte("http://example.com/3"),
	}
	assert.Equal(test, wantAttributes, builder.attributes)
}
