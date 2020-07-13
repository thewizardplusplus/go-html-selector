package htmlselector

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
