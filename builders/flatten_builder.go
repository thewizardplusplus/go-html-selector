package builders

import (
	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
)

// FlattenBuilder ...
type FlattenBuilder struct {
	attributeValues [][]byte
}

// AttributeValues ...
func (builder FlattenBuilder) AttributeValues() [][]byte {
	return builder.attributeValues
}

// AddTag ...
func (builder FlattenBuilder) AddTag(name []byte) {}

// AddAttribute ...
func (builder *FlattenBuilder) AddAttribute(name []byte, value []byte) {
	valueCopy := byteutils.Copy(value)
	builder.attributeValues = append(builder.attributeValues, valueCopy)
}
