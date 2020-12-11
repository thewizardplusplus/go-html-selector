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
	config := newOptionConfig(options)

	selector := newSelector(reader, builder)
	universalTagAttributeFilters := filters[UniversalTag]
	// use the special form of a type assertion to avoid panic; a nil result
	// is processed separately below
	textBuilder, _ := builder.(TextBuilder)
	for {
		switch selector.tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			selector.selectTag(filters, universalTagAttributeFilters, config)
		case html.TextToken:
			selector.selectText(textBuilder, config)
		case html.ErrorToken:
			if err := selector.tokenizer.Err(); err != io.EOF {
				return err
			}

			return nil
		}
	}
}
