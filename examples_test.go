package htmlselector_test

import (
	"fmt"
	"log"
	"strings"

	htmlselector "github.com/thewizardplusplus/go-html-selector"
	"github.com/thewizardplusplus/go-html-selector/builders"
)

func ExampleSelectTags_withStructuralBuilder() {
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
			<li>
				<a>3</a>
				<video></video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

	var builder builders.StructuralBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		&builder,
		htmlselector.SkipEmptyTags(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, tag := range builder.Tags() {
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

func ExampleSelectTags_withFlattenBuilder() {
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
			<li>
				<a>3</a>
				<video></video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

	var builder builders.FlattenBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		&builder,
		htmlselector.SkipEmptyTags(),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, attributeValue := range builder.AttributeValues() {
		fmt.Printf("%q\n", attributeValue)
	}

	// Output:
	// "http://example.com/1"
	// "http://example.com/1.1"
	// "http://example.com/1.2"
	// "http://example.com/2"
	// "http://example.com/2.1"
	// "http://example.com/2.2"
}
