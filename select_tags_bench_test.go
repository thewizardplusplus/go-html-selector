package htmlselector

import (
	"fmt"
	"strings"
	"testing"

	"code.cloudfoundry.org/bytefmt"
	"github.com/thewizardplusplus/go-html-selector/builders"
)

func BenchmarkSelectTags(benchmark *testing.B) {
	type args struct {
		makeMarkup func(tagCount int) string
		filters    OptimizedFilterGroup
	}

	for _, data := range []struct {
		name string
		args args
	}{
		{
			name: "simple markup",
			args: args{
				makeMarkup: func(tagCount int) string {
					var markupParts []string
					for tagIndex := 0; tagIndex < tagCount; tagIndex++ {
						markupParts = append(markupParts, fmt.Sprintf(
							`<p><a href="http://example.com/%[1]d">%[1]d</a></p>`,
							tagIndex,
						))
					}

					return strings.Join(markupParts, "")
				},
				filters: OptimizedFilterGroup{"a": {"href": {}}},
			},
		},
		{
			name: "complex markup",
			args: args{
				makeMarkup: func(tagCount int) string {
					var markupParts []string
					for tagIndex := 0; tagIndex < tagCount; tagIndex++ {
						markupParts = append(markupParts, fmt.Sprintf(
							`<p><a href="http://example.com/%[1]d" title="%[1]d">%[1]d</a></p>`+
								`<p><img src="http://example.com/%[1]d" alt="%[1]d" /></p>`,
							tagIndex,
						))
					}

					return strings.Join(markupParts, "")
				},
				filters: OptimizedFilterGroup{
					"a":   {"href": {}, "title": {}},
					"img": {"src": {}, "alt": {}},
				},
			},
		},
	} {
		for tagCount := 10; tagCount <= 1e6; tagCount *= 10 {
			markup := data.args.makeMarkup(tagCount)
			markupSize := bytefmt.ByteSize(uint64(len(markup)))

			name := fmt.Sprintf("%s/%d tags/%s", data.name, tagCount, markupSize)
			benchmark.Run(name, func(benchmark *testing.B) {
				reader := strings.NewReader(markup)
				benchmark.ResetTimer()

				for i := 0; i < benchmark.N; i++ {
					var builder builders.StructuralBuilder
					SelectTags(reader, data.args.filters, &builder) // nolint: errcheck
				}
			})
		}
	}
}
