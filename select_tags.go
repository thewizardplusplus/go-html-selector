package htmlselector

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

// Filter ...
type Filter struct {
	Tag        []byte
	Attributes [][]byte
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
func SelectTags(reader io.Reader, filters []Filter) ([]Tag, error) {
	var tags []Tag
	tokenizer := html.NewTokenizer(reader)
	for {
		switch tokenizer.Next() {
		case html.StartTagToken, html.SelfClosingTagToken:
			tagName, hasAttribute := tokenizer.TagName()
			matchedFilterIndex := -1
			for filterIndex, filter := range filters {
				if bytes.Equal(filter.Tag, tagName) {
					matchedFilterIndex = filterIndex
					break
				}
			}
			if matchedFilterIndex == -1 {
				continue
			}

			var attributes []Attribute
			for hasAttribute {
				var attributeName, attributeValue []byte
				var attributeMatched bool
				attributeName, attributeValue, hasAttribute = tokenizer.TagAttr()
				for _, attribute := range filters[matchedFilterIndex].Attributes {
					if bytes.Equal(attribute, attributeName) {
						attributeMatched = true
						break
					}
				}
				if !attributeMatched {
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
