package builders

import (
	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
)

// FlattenBuilder ...
type FlattenBuilder struct {
	attributes [][]byte
}

// Attributes ...
func (builder FlattenBuilder) Attributes() [][]byte {
	return builder.attributes
}

// AddTag ...
func (builder FlattenBuilder) AddTag(name []byte) {}

// AddAttribute ...
func (builder *FlattenBuilder) AddAttribute(name []byte, value []byte) {
	valueCopy := byteutils.Copy(value)
	builder.attributes = append(builder.attributes, valueCopy)
}
