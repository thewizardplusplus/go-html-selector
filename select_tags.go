package htmlselector

import (
	"io"

	"golang.org/x/net/html"
)

// SelectTags ...
func SelectTags(
	reader io.Reader,
	filters OptimizedFilterGroup,
	builder Builder,
	options ...Option,
) error {
	selector := newSelector(reader, builder, options...)
	universalTagAttributeFilters := filters[UniversalTag]
	// use the special form of a type assertion to avoid panic; a nil result
	// is processed separately below
	selectionTerminator, _ := builder.(SelectionTerminator)
	for {
		switch selector.nextToken() {
		case html.StartTagToken, html.SelfClosingTagToken:
			selector.selectTag(filters, universalTagAttributeFilters)
		case html.TextToken:
			selector.selectText()
		case html.ErrorToken:
			return selector.error()
		}

		if selectionTerminator != nil && selectionTerminator.IsSelectionTerminated() {
			return nil
		}
	}
}
