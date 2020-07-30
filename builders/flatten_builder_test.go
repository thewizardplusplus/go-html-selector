package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFlattenBuilder_Attributes(test *testing.T) {
	builder := FlattenBuilder{
		attributeValues: [][]byte{
			[]byte("http://example.com/1"),
			[]byte("http://example.com/2"),
		},
	}
	gotAttributeValues := builder.AttributeValues()

	assert.Equal(test, builder.attributeValues, gotAttributeValues)
}

func TestFlattenBuilder_AddTag(test *testing.T) {
	builder := FlattenBuilder{
		attributeValues: [][]byte{
			[]byte("http://example.com/1"),
			[]byte("http://example.com/2"),
		},
	}
	builder.AddTag([]byte("a"))

	wantAttributeValues := [][]byte{
		[]byte("http://example.com/1"),
		[]byte("http://example.com/2"),
	}
	assert.Equal(test, wantAttributeValues, builder.attributeValues)
}

func TestFlattenBuilder_AddAttribute(test *testing.T) {
	builder := FlattenBuilder{
		attributeValues: [][]byte{
			[]byte("http://example.com/1"),
			[]byte("http://example.com/2"),
		},
	}
	builder.AddAttribute([]byte("href"), []byte("http://example.com/3"))

	wantAttributeValues := [][]byte{
		[]byte("http://example.com/1"),
		[]byte("http://example.com/2"),
		[]byte("http://example.com/3"),
	}
	assert.Equal(test, wantAttributeValues, builder.attributeValues)
}
