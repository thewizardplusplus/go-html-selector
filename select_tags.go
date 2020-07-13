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
			tagName, _ := tokenizer.TagName()
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

			tagNameCopy := make([]byte, len(tagName))
			copy(tagNameCopy, tagName)

			tags = append(tags, Tag{Name: tagNameCopy})
		case html.ErrorToken:
			if err := tokenizer.Err(); err != io.EOF {
				return nil, err
			}

			return tags, nil
		}
	}
}
