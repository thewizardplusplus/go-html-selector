package htmlselector

import (
	"io"

	"golang.org/x/net/html"
)

//go:generate mockery -name=Builder -inpkg -case=underscore -testonly

// Builder ...
type Builder interface {
	AddTag(name []byte)
	AddAttribute(name []byte, value []byte)
}

//go:generate mockery -name=TextBuilder -inpkg -case=underscore -testonly

// TextBuilder ...
type TextBuilder interface {
	AddText(text []byte)
}

// MultiBuilder ...
type MultiBuilder struct {
	Builder
	TextBuilder
}

// SelectTags ...
func SelectTags(
	reader io.Reader,
	filters OptimizedFilterGroup,
	builder Builder,
	options ...Option,
) error {
	selector := newSelector(reader, builder, options...)
	universalTagAttributeFilters := filters[UniversalTag]
	for {
		switch selector.tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			selector.selectTag(filters, universalTagAttributeFilters)
		case html.TextToken:
			selector.selectText()
		case html.ErrorToken:
			if err := selector.tokenizer.Err(); err != io.EOF {
				return err
			}

			return nil
		}
	}
}
