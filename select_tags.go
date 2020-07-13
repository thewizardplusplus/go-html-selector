package htmlselector

import (
	"io"
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
	return tags, nil
}
