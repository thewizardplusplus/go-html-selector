package htmlselector

import (
	"fmt"
	"strings"
	"testing"

	"code.cloudfoundry.org/bytefmt"
)

func BenchmarkDistance(benchmark *testing.B) {
	type args struct {
		makeMarkup func(tagCount int) string
		filters    []Filter
	}

	for _, data := range []struct {
		name string
		args args
	}{
		{
			name: "simple markup",
			args: args{
				makeMarkup: func(tagCount int) string {
					var tags []string
					for index := 0; index < tagCount; index++ {
						tags = append(tags, fmt.Sprintf(
							`  <li><a href="http://example.com/%[1]d">%[1]d</a></li>`,
							index,
						))
					}

					return fmt.Sprintf("<ul>\n%s\n</ul>", strings.Join(tags, "\n"))
				},
				filters: []Filter{{[]byte("a"), [][]byte{[]byte("href")}}},
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
					SelectTags(reader, data.args.filters) // nolint: errcheck
				}
			})
		}
	}
}
