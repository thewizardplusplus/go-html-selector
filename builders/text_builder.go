package builders

// TextBuilder ...
type TextBuilder struct {
	textParts [][]byte
}

// TextParts ...
func (builder TextBuilder) TextParts() [][]byte {
	return builder.textParts
}
