package htmlselector_test

import (
	"bytes"
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
				<a href="">3</a>
				<video src="" poster=""></video>
			</li>
			<li>
				<a>4</a>
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
		htmlselector.SkipEmptyAttributes(),
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
				<a href="">3</a>
				<video src="" poster=""></video>
			</li>
			<li>
				<a>4</a>
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
		htmlselector.SkipEmptyAttributes(),
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

func ExampleSelectTags_withUniversalTag() {
	reader := strings.NewReader(`
		<ul>
			<li>
				<a href="http://example.com/1" title="link #1">1</a>
				<video
					src="http://example.com/1.1"
					poster="http://example.com/1.2"
					title="video #1">
				</video>
			</li>
			<li>
				<a href="http://example.com/2" title="link #2">2</a>
				<video
					src="http://example.com/2.1"
					poster="http://example.com/2.2"
					title="video #2">
				</video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		htmlselector.UniversalTag: {"title", "href"},
		"video":                   {"src", "poster"},
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
	//   title="link #1"
	// <video>:
	//   src="http://example.com/1.1"
	//   poster="http://example.com/1.2"
	//   title="video #1"
	// <a>:
	//   href="http://example.com/2"
	//   title="link #2"
	// <video>:
	//   src="http://example.com/2.1"
	//   poster="http://example.com/2.2"
	//   title="video #2"
}

func ExampleSelectTags_withTextBuilder() {
	reader := strings.NewReader(`
		<ul>
			<li>
				link #1: <a href="http://example.com/1">one</a>
				<video
					src="http://example.com/1.1"
					poster="http://example.com/1.2">
					Unable to embed video #1.
				</video>
			</li>
			<li>
				link #2: <a href="http://example.com/2">two</a>
				<video
					src="http://example.com/2.1"
					poster="http://example.com/2.2">
					Unable to embed video #2.
				</video>
			</li>
		</ul>
	`)

	filters := htmlselector.OptimizeFilters(htmlselector.FilterGroup{
		"a":     {"href"},
		"video": {"src", "poster"},
	})

	var builder builders.StructuralBuilder
	var textBuilder builders.TextBuilder
	err := htmlselector.SelectTags(
		reader,
		filters,
		htmlselector.MultiBuilder{Builder: &builder, TextBuilder: &textBuilder},
		htmlselector.SkipEmptyText(),
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

	fmt.Println("text parts:")
	for _, textPart := range textBuilder.TextParts() {
		textPart = bytes.TrimSpace(textPart)
		fmt.Printf("  %q\n", textPart)
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
	// text parts:
	//   "link #1:"
	//   "one"
	//   "Unable to embed video #1."
	//   "link #2:"
	//   "two"
	//   "Unable to embed video #2."
}
