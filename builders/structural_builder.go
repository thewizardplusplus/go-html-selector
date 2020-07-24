package builders

import (
	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
)

// Tag ...
type Tag struct {
	Name       []byte
	Attributes []Attribute
}

// Attribute ...
type Attribute struct {
	Name  []byte
	Value []byte
}

// StructuralBuilder ...
type StructuralBuilder struct {
	tags       []Tag
	attributes []Attribute
}

// Tags ...
func (builder StructuralBuilder) Tags() []Tag {
	return builder.tags
}

// AddTag ...
func (builder *StructuralBuilder) AddTag(name []byte) {
	builder.tags = append(builder.tags, Tag{
		Name:       byteutils.Copy(name),
		Attributes: builder.attributes,
	})

	builder.attributes = nil
}

// AddAttribute ...
func (builder *StructuralBuilder) AddAttribute(name []byte, value []byte) {
	builder.attributes = append(builder.attributes, Attribute{
		Name:  byteutils.Copy(name),
		Value: byteutils.Copy(value),
	})
}
