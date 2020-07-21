package htmlselector

import (
	"io"

	"golang.org/x/net/html"
)

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
	options ...Option,
) ([]Tag, error) {
	config := newOptionConfig(options)

	var tags []Tag
	tokenizer := html.NewTokenizer(reader)
	for {
		switch tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			tagName, hasAttribute := tokenizer.TagName()
			attributeFilters, ok := filters[TagName(bytesToString(tagName))]
			if !ok {
				continue
			}

			var attributes []Attribute
			for hasAttribute {
				var attributeName, attributeValue []byte
				attributeName, attributeValue, hasAttribute = tokenizer.TagAttr()
				_, ok := attributeFilters[AttributeName(bytesToString(attributeName))]
				if !ok {
					continue
				}

				attributes = append(attributes, Attribute{
					Name:  copyBytes(attributeName),
					Value: copyBytes(attributeValue),
				})
			}
			if config.skipEmptyTags && len(attributes) == 0 {
				continue
			}

			tags = append(tags, Tag{
				Name:       copyBytes(tagName),
				Attributes: attributes,
			})
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				return nil, err
			}

			return tags, nil
		}
	}
}
