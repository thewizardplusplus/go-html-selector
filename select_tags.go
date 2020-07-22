package htmlselector

import (
	"io"

	"golang.org/x/net/html"
)

//go:generate mockery -name=Builder -inpkg -case=underscore -testonly

// Builder ...
type Builder interface {
	StartTag(name []byte)
	FinishTag()
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
			tagName, hasAttribute := tokenizer.TagName()
			attributeFilters, ok := filters[TagName(bytesToString(tagName))]
			if !ok {
				continue
			}

			builder.StartTag(tagName)

			var attributeCount int
			for hasAttribute {
				var attributeName, attributeValue []byte
				attributeName, attributeValue, hasAttribute = tokenizer.TagAttr()
				_, ok := attributeFilters[AttributeName(bytesToString(attributeName))]
				if !ok {
					continue
				}

				builder.AddAttribute(attributeName, attributeValue)
				attributeCount++
			}
			if config.skipEmptyTags && attributeCount == 0 {
				continue
			}

			builder.FinishTag()
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				return err
			}

			return nil
		}
	}
}
