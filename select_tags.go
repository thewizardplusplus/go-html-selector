package htmlselector

import (
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

// SelectTags ...
func SelectTags(
	reader io.Reader,
	filters OptimizedFilterGroup,
	builder Builder,
	options ...Option,
) error {
	config := newOptionConfig(options)

	tokenizer := html.NewTokenizer(reader)
	for {
		switch tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			selectTag(tokenizer, filters, builder, config.skipEmptyTags)
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
	builder Builder,
	skipEmptyTags bool,
) {
	name, _ := tokenizer.TagName()
	attributeFilters, ok := filters[TagName(byteutils.String(name))]
	if !ok {
		return
	}

	attributeCount := selectAttributes(tokenizer, attributeFilters, builder)
	if skipEmptyTags && attributeCount == 0 {
		return
	}

	builder.AddTag(name)
}

func selectAttributes(
	tokenizer *html.Tokenizer,
	filters OptimizedAttributeFilterGroup,
	builder Builder,
) (count int) {
	hasNext := true
	for hasNext {
		var name, value []byte
		name, value, hasNext = tokenizer.TagAttr()
		if _, ok := filters[AttributeName(byteutils.String(name))]; !ok {
			continue
		}

		builder.AddAttribute(name, value)
		count++
	}

	return count
}
