package htmlselector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimizeFilters(test *testing.T) {
	type args struct {
		filters FilterGroup
	}

	for _, data := range []struct {
		name string
		args args
		want OptimizedFilterGroup
	}{
		{
			name: "without filters",
			args: args{
				filters: nil,
			},
			want: OptimizedFilterGroup{},
		},
		{
			name: "with an empty filter",
			args: args{
				filters: FilterGroup{"a": nil},
			},
			want: OptimizedFilterGroup{"a": {}},
		},
		{
			name: "with a nonempty filter",
			args: args{
				filters: FilterGroup{"a": {"href", "title"}},
			},
			want: OptimizedFilterGroup{"a": {"href": {}, "title": {}}},
		},
		{
			name: "with few filters",
			args: args{
				filters: FilterGroup{"a": {"href", "title"}, "img": {"src", "alt"}},
			},
			want: OptimizedFilterGroup{
				"a":   {"href": {}, "title": {}},
				"img": {"src": {}, "alt": {}},
			},
		},
	} {
		test.Run(data.name, func(t *testing.T) {
			got := OptimizeFilters(data.args.filters)

			assert.Equal(test, data.want, got)
		})
	}
}
