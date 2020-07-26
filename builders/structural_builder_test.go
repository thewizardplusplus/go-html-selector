package builders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructuralBuilder_Tags(test *testing.T) {
	builder := StructuralBuilder{
		tags: []Tag{
			{
				Name:       []byte("a"),
				Attributes: []Attribute{{[]byte("href"), []byte("http://example.com/1")}},
			},
			{
				Name:       []byte("img"),
				Attributes: []Attribute{{[]byte("src"), []byte("http://example.com/2")}},
			},
		},
	}
	gotTags := builder.Tags()

	assert.Equal(test, builder.tags, gotTags)
}

func TestStructuralBuilder_AddTag(test *testing.T) {
	builder := StructuralBuilder{
		tags: []Tag{
			{
				Name:       []byte("a"),
				Attributes: []Attribute{{[]byte("href"), []byte("http://example.com/1")}},
			},
			{
				Name:       []byte("img"),
				Attributes: []Attribute{{[]byte("src"), []byte("http://example.com/2")}},
			},
		},
		attributes: []Attribute{
			{[]byte("src"), []byte("http://example.com/3")},
			{[]byte("poster"), []byte("http://example.com/3.1")},
		},
	}
	builder.AddTag([]byte("video"))

	wantTags := []Tag{
		{
			Name:       []byte("a"),
			Attributes: []Attribute{{[]byte("href"), []byte("http://example.com/1")}},
		},
		{
			Name:       []byte("img"),
			Attributes: []Attribute{{[]byte("src"), []byte("http://example.com/2")}},
		},
		{
			Name: []byte("video"),
			Attributes: []Attribute{
				{[]byte("src"), []byte("http://example.com/3")},
				{[]byte("poster"), []byte("http://example.com/3.1")},
			},
		},
	}
	assert.Equal(test, wantTags, builder.tags)
	assert.Nil(test, builder.attributes)
}

func TestStructuralBuilder_AddAttribute(test *testing.T) {
	builder := StructuralBuilder{
		attributes: []Attribute{
			{[]byte("src"), []byte("http://example.com/1")},
			{[]byte("poster"), []byte("http://example.com/1.1")},
		},
	}
	builder.AddAttribute([]byte("title"), []byte("1"))

	wantAttributes := []Attribute{
		{[]byte("src"), []byte("http://example.com/1")},
		{[]byte("poster"), []byte("http://example.com/1.1")},
		{[]byte("title"), []byte("1")},
	}
	assert.Equal(test, wantAttributes, builder.attributes)
}
