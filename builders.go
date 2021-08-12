package htmlselector

//go:generate mockery --name=Builder --inpackage --case=underscore --testonly

// Builder ...
type Builder interface {
	AddTag(name []byte)
	AddAttribute(name []byte, value []byte)
}

//go:generate mockery --name=TextBuilder --inpackage --case=underscore --testonly

// TextBuilder ...
type TextBuilder interface {
	AddText(text []byte)
}

// MultiBuilder ...
type MultiBuilder struct {
	Builder
	TextBuilder
}

// SelectionTerminator ...
type SelectionTerminator interface {
	IsSelectionTerminated() bool
}
