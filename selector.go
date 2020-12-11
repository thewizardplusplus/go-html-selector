package htmlselector

import (
	"io"

	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
	"golang.org/x/net/html"
)

type selector struct {
	tokenizer *html.Tokenizer
}

func newSelector(reader io.Reader) selector {
	tokenizer := html.NewTokenizer(reader)
	return selector{
		tokenizer: tokenizer,
	}
}

func (selector selector) selectTag(
	filters OptimizedFilterGroup,
	additionalAttributeFilters OptimizedAttributeFilterGroup,
	builder Builder,
	config OptionConfig,
) {
	name, hasAttributes := selector.tokenizer.TagName()
	attributeFilters, ok := filters[TagName(byteutils.String(name))]
	if !ok && len(additionalAttributeFilters) == 0 {
		return
	}

	attributeCount := selectAttributes(
		selector.tokenizer,
		hasAttributes,
		attributeFilters,
		additionalAttributeFilters,
		builder,
		config,
	)
	if config.skipEmptyTags && attributeCount == 0 {
		return
	}

	builder.AddTag(name)
}
