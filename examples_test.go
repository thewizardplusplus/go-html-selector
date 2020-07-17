package htmlselector_test

import (
	"fmt"
	"log"
	"strings"

	htmlselector "github.com/thewizardplusplus/go-html-selector"
)

func ExampleSelectTags() {
	reader := strings.NewReader(`
		<ul>
			<li>
				<a href="http://example.com/1">1</a>
				<video
					src="http://example.com/1.1"
					poster="http://example.com/1.2">
				</video>
			</li>
			<li>
				<a href="http://example.com/2">2</a>
				<video
					src="http://example.com/2.1"
					poster="http://example.com/2.2">
				</video>
			</li>
		</ul>
	`)

	filters := htmlselector.FilterGroup{
		"a":     {"href": {}},
		"video": {"src": {}, "poster": {}},
	}

	tags, err := htmlselector.SelectTags(reader, filters)
	if err != nil {
		log.Fatal(err)
	}

	for _, tag := range tags {
		fmt.Printf("<%s>:\n", tag.Name)
		for _, attribute := range tag.Attributes {
			fmt.Printf("  %s=%q\n", attribute.Name, attribute.Value)
		}
	}

	// Output:
	// <a>:
	//   href="http://example.com/1"
	// <video>:
	//   src="http://example.com/1.1"
	//   poster="http://example.com/1.2"
	// <a>:
	//   href="http://example.com/2"
	// <video>:
	//   src="http://example.com/2.1"
	//   poster="http://example.com/2.2"
}
