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
	for {
		switch selector.nextToken() {
		case html.StartTagToken, html.SelfClosingTagToken:
			selector.selectTag(filters, universalTagAttributeFilters)
		case html.TextToken:
			selector.selectText()
		case html.ErrorToken:
			return selector.error()
		}
	}
}
