package htmlselector

import (
	"io"

	"golang.org/x/net/html"
)

type selector struct {
	tokenizer *html.Tokenizer
}

func newSelector(reader io.Reader) selector {
	tokenizer := html.NewTokenizer(reader)
	return selector{
		tokenizer: tokenizer,
	}
}
