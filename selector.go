package htmlselector

import (
	"bytes"
	"io"

	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
	"golang.org/x/net/html"
)

type selector struct {
	tokenizer   *html.Tokenizer
	builder     Builder
	textBuilder TextBuilder
}

func newSelector(reader io.Reader, builder Builder) selector {
	tokenizer := html.NewTokenizer(reader)
	// use the special form of a type assertion to avoid panic; a nil result
	// is processed separately below
	textBuilder, _ := builder.(TextBuilder)
	return selector{
		tokenizer:   tokenizer,
		builder:     builder,
		textBuilder: textBuilder,
	}
}

func (selector selector) selectTag(
	filters OptimizedFilterGroup,
	additionalAttributeFilters OptimizedAttributeFilterGroup,
	config OptionConfig,
) {
	name, hasAttributes := selector.tokenizer.TagName()
	attributeFilters, ok := filters[TagName(byteutils.String(name))]
	if !ok && len(additionalAttributeFilters) == 0 {
		return
	}

	attributeCount := selector.selectAttributes(
		hasAttributes,
		attributeFilters,
		additionalAttributeFilters,
		config,
	)
	if config.skipEmptyTags && attributeCount == 0 {
		return
	}

	selector.builder.AddTag(name)
}

func (selector selector) selectAttributes(
	hasAttributes bool,
	filters OptimizedAttributeFilterGroup,
	additionalFilters OptimizedAttributeFilterGroup,
	config OptionConfig,
) (count int) {
	hasNext := hasAttributes
	for hasNext {
		var name, value []byte
		name, value, hasNext = selector.tokenizer.TagAttr()
		filterName := AttributeName(byteutils.String(name))
		if _, ok := filters[filterName]; !ok {
			if _, ok := additionalFilters[filterName]; !ok {
				continue
			}
		}
		if config.skipEmptyAttributes && len(value) == 0 {
			continue
		}

		selector.builder.AddAttribute(name, value)
		count++
	}

	return count
}

func (selector selector) selectText(config OptionConfig) {
	if selector.textBuilder == nil {
		return
	}

	text := selector.tokenizer.Raw()
	// bytes.TrimSpace doesn't make new allocations and also has the optimization
	// for an ASCII-only text, so it's optimal to use it
	if config.skipEmptyText && len(bytes.TrimSpace(text)) == 0 {
		return
	}

	selector.textBuilder.AddText(text)
}
