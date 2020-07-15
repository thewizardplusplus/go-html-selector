package htmlselector

import (
	"fmt"
	"strings"
	"testing"

	"code.cloudfoundry.org/bytefmt"
)

func BenchmarkDistance(benchmark *testing.B) {
	for tagCount := 10; tagCount <= 1e6; tagCount *= 10 {
		markup := generateMarkup(tagCount)
		name :=
			fmt.Sprintf("%dTags/%s", tagCount, bytefmt.ByteSize(uint64(len(markup))))
		benchmark.Run(name, func(benchmark *testing.B) {
			reader := strings.NewReader(markup)
			filters := []Filter{{[]byte("a"), [][]byte{[]byte("href")}}}
			benchmark.ResetTimer()

			for i := 0; i < benchmark.N; i++ {
				SelectTags(reader, filters) // nolint: errcheck
			}
		})
	}
}

func generateMarkup(tagCount int) string {
	var tags []string
	for index := 0; index < tagCount; index++ {
		tags = append(tags, fmt.Sprintf(
			`  <li><a href="http://example.com/%[1]d">%[1]d</a></li>`,
			index,
		))
	}

	return fmt.Sprintf("<ul>\n%s\n</ul>", strings.Join(tags, "\n"))
}
