package htmlselector

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterGroup_Unmarshal(test *testing.T) {
	type args struct {
		data []byte
	}

	for _, data := range []struct {
		name string
		args args
		want FilterGroup
	}{
		{
			name: "without filters",
			args: args{
				data: []byte("null"),
			},
			want: nil,
		},
		{
			name: "with an empty filter",
			args: args{
				data: []byte(`{"a": null}`),
			},
			want: FilterGroup{"a": nil},
		},
		{
			name: "with a nonempty filter",
			args: args{
				data: []byte(`{"a": ["href", "title"]}`),
			},
			want: FilterGroup{"a": {"href", "title"}},
		},
		{
			name: "with few filters",
			args: args{
				data: []byte(`{"a": ["href", "title"], "img": ["src", "alt"]}`),
			},
			want: FilterGroup{"a": {"href", "title"}, "img": {"src", "alt"}},
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			var filters FilterGroup
			err := json.Unmarshal(data.args.data, &filters)

			assert.Equal(test, data.want, filters)
			assert.NoError(test, err)
		})
	}
}

func TestOptimizeFilters(test *testing.T) {
	type args struct {
		filters FilterGroup
		options []Option
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
				options: nil,
			},
			want: OptimizedFilterGroup{},
		},
		{
			name: "with an empty filter/without skipping",
			args: args{
				filters: FilterGroup{"a": nil},
				options: nil,
			},
			want: OptimizedFilterGroup{"a": {}},
		},
		{
			name: "with an empty filter/with skipping",
			args: args{
				filters: FilterGroup{"a": nil},
				options: []Option{SkipEmptyTags()},
			},
			want: OptimizedFilterGroup{},
		},
		{
			name: "with a nonempty filter",
			args: args{
				filters: FilterGroup{"a": {"href", "title"}},
				options: nil,
			},
			want: OptimizedFilterGroup{"a": {"href": {}, "title": {}}},
		},
		{
			name: "with few filters",
			args: args{
				filters: FilterGroup{"a": {"href", "title"}, "img": {"src", "alt"}},
				options: nil,
			},
			want: OptimizedFilterGroup{
				"a":   {"href": {}, "title": {}},
				"img": {"src": {}, "alt": {}},
			},
		},
		{
			name: "with the universal tag",
			args: args{
				filters: FilterGroup{
					UniversalTag: {"title", "alt"},
					"a":          {"href", "title"},
					"img":        {"src", "alt"},
				},
				options: nil,
			},
			want: OptimizedFilterGroup{
				UniversalTag: {"title": {}, "alt": {}},
				"a":          {"href": {}},
				"img":        {"src": {}},
			},
		},
	} {
		test.Run(data.name, func(t *testing.T) {
			got := OptimizeFilters(data.args.filters, data.args.options...)

			assert.Equal(test, data.want, got)
		})
	}
}
