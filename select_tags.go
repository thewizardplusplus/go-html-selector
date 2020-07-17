package htmlselector

import (
	"io"

	"golang.org/x/net/html"
)

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
func SelectTags(reader io.Reader, filters FilterGroup) ([]Tag, error) {
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
