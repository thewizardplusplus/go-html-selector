package htmlselector

import (
	"bytes"
	"io"

	byteutils "github.com/thewizardplusplus/go-html-selector/byte-utils"
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

	tokenizer := html.NewTokenizer(reader)
	universalTagAttributeFilters := filters[UniversalTag]
	// use the special form of a type assertion to avoid panic; a nil result
	// is processed separately below
	textBuilder, _ := builder.(TextBuilder)
	for {
		switch tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			selectTag(tokenizer, filters, universalTagAttributeFilters, builder, config)
		case html.TextToken:
			selectText(tokenizer, textBuilder, config)
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				return err
			}

			return nil
		}
	}
}

func selectTag(
	tokenizer *html.Tokenizer,
	filters OptimizedFilterGroup,
	additionalAttributeFilters OptimizedAttributeFilterGroup,
	builder Builder,
	config OptionConfig,
) {
	name, hasAttributes := tokenizer.TagName()
	attributeFilters, ok := filters[TagName(byteutils.String(name))]
	if !ok && len(additionalAttributeFilters) == 0 {
		return
	}

	attributeCount := selectAttributes(
		tokenizer,
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

func selectAttributes(
	tokenizer *html.Tokenizer,
	hasAttributes bool,
	filters OptimizedAttributeFilterGroup,
	additionalFilters OptimizedAttributeFilterGroup,
	builder Builder,
	config OptionConfig,
) (count int) {
	hasNext := hasAttributes
	for hasNext {
		var name, value []byte
		name, value, hasNext = tokenizer.TagAttr()
		filterName := AttributeName(byteutils.String(name))
		if _, ok := filters[filterName]; !ok {
			if _, ok := additionalFilters[filterName]; !ok {
				continue
			}
		}
		if config.skipEmptyAttributes && len(value) == 0 {
			continue
		}

		builder.AddAttribute(name, value)
		count++
	}

	return count
}

func selectText(
	tokenizer *html.Tokenizer,
	textBuilder TextBuilder,
	config OptionConfig,
) {
	if textBuilder == nil {
		return
	}

	text := tokenizer.Raw()
	// bytes.TrimSpace doesn't make new allocations and also has the optimization
	// for an ASCII-only text, so it's optimal to use it
	if config.skipEmptyText && len(bytes.TrimSpace(text)) == 0 {
		return
	}

	textBuilder.AddText(text)
}
