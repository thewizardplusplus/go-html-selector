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

// Attribute ...
type Attribute struct {
	Name  []byte
	Value []byte
}

// Tag ...
type Tag struct {
	Name       []byte
	Attributes []Attribute
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
			tagName, _ := tokenizer.TagName()
			attributeFilters, ok := filters[TagName(bytesToString(tagName))]
			if !ok {
				continue
			}

			attributeCount := selectAttributes(tokenizer, attributeFilters, builder)
			if config.skipEmptyTags && attributeCount == 0 {
				continue
			}

			builder.AddTag(tagName)
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				return err
			}

			return nil
		}
	}
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
		if _, ok := filters[AttributeName(bytesToString(name))]; !ok {
			continue
		}

		builder.AddAttribute(name, value)
		count++
	}

	return count
}
