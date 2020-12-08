package builders

import (
	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
)

// TextBuilder ...
type TextBuilder struct {
	textParts [][]byte
}

// TextParts ...
func (builder TextBuilder) TextParts() [][]byte {
	return builder.textParts
}

// AddText ...
func (builder *TextBuilder) AddText(text []byte) {
	textCopy := byteutils.Copy(text)
	builder.textParts = append(builder.textParts, textCopy)
}
