package htmlselector

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
)

func TestSelectTags(test *testing.T) {
	type args struct {
		reader  io.Reader
		filters OptimizedFilterGroup
		options []Option
	}

	for _, data := range []struct {
		name     string
		args     args
		wantTags []Tag
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			name: "success/with empty arguments",
			args: args{
				reader:  strings.NewReader(""),
				filters: nil,
				options: nil,
			},
			wantTags: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "success/with an empty reader",
			args: args{
				reader:  strings.NewReader(""),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				options: nil,
			},
			wantTags: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "success/without filters",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><a href="http://example.com/1">1</a></li>
						<li><a href="http://example.com/2">2</a></li>
					</ul>
				`),
				filters: nil,
				options: nil,
			},
			wantTags: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "success/with a conventional tag",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><a href="http://example.com/1">1</a></li>
						<li><a href="http://example.com/2">2</a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("a"),
					Attributes: []Attribute{
						{
							Name:  []byte("href"),
							Value: []byte("http://example.com/1"),
						},
					},
				},
				{
					Name: []byte("a"),
					Attributes: []Attribute{
						{
							Name:  []byte("href"),
							Value: []byte("http://example.com/2"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with a self-closing tag",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><img src="http://example.com/1" /></li>
						<li><img src="http://example.com/2" /></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"img": {"src": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("img"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/1"),
						},
					},
				},
				{
					Name: []byte("img"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/2"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with few tags",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<a href="http://example.com/1"><img src="http://example.com/1.1" /></a>
						</li>
						<li>
							<a href="http://example.com/2"><img src="http://example.com/2.1" /></a>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}, "img": {"src": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("a"),
					Attributes: []Attribute{
						{
							Name:  []byte("href"),
							Value: []byte("http://example.com/1"),
						},
					},
				},
				{
					Name: []byte("img"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/1.1"),
						},
					},
				},
				{
					Name: []byte("a"),
					Attributes: []Attribute{
						{
							Name:  []byte("href"),
							Value: []byte("http://example.com/2"),
						},
					},
				},
				{
					Name: []byte("img"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/2.1"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with missed tags",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><a href="http://example.com/1">1</a></li>
						<li><a href="http://example.com/2">2</a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}, "img": {"src": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("a"),
					Attributes: []Attribute{
						{
							Name:  []byte("href"),
							Value: []byte("http://example.com/1"),
						},
					},
				},
				{
					Name: []byte("a"),
					Attributes: []Attribute{
						{
							Name:  []byte("href"),
							Value: []byte("http://example.com/2"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/without attributes/without skipping/by markup and filters",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><video></video></li>
						<li><video></video></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": nil},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name:       []byte("video"),
					Attributes: nil,
				},
				{
					Name:       []byte("video"),
					Attributes: nil,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/without attributes/without skipping/by markup",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><video></video></li>
						<li><video></video></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": {"src": {}, "poster": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name:       []byte("video"),
					Attributes: nil,
				},
				{
					Name:       []byte("video"),
					Attributes: nil,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/without attributes/without skipping/by filters",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<video
								src="http://example.com/1"
								poster="http://example.com/1.1">
							</video>
						</li>
						<li>
							<video
								src="http://example.com/2"
								poster="http://example.com/2.1">
							</video>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": nil},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name:       []byte("video"),
					Attributes: nil,
				},
				{
					Name:       []byte("video"),
					Attributes: nil,
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/without attributes/with skipping/by markup and filters",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><video></video></li>
						<li><video></video></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": nil},
				options: []Option{SkipEmptyTags()},
			},
			wantTags: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "success/without attributes/with skipping/by markup",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><video></video></li>
						<li><video></video></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": {"src": {}, "poster": {}}},
				options: []Option{SkipEmptyTags()},
			},
			wantTags: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "success/without attributes/with skipping/by filters",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<video
								src="http://example.com/1"
								poster="http://example.com/1.1">
							</video>
						</li>
						<li>
							<video
								src="http://example.com/2"
								poster="http://example.com/2.1">
							</video>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": nil},
				options: []Option{SkipEmptyTags()},
			},
			wantTags: nil,
			wantErr:  assert.NoError,
		},
		{
			name: "success/with few attributes",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<video
								src="http://example.com/1"
								poster="http://example.com/1.1">
							</video>
						</li>
						<li>
							<video
								src="http://example.com/2"
								poster="http://example.com/2.1">
							</video>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": {"src": {}, "poster": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("video"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/1"),
						},
						{
							Name:  []byte("poster"),
							Value: []byte("http://example.com/1.1"),
						},
					},
				},
				{
					Name: []byte("video"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/2"),
						},
						{
							Name:  []byte("poster"),
							Value: []byte("http://example.com/2.1"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with missed attributes",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<video src="http://example.com/1"></video>
						</li>
						<li>
							<video src="http://example.com/2"></video>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": {"src": {}, "poster": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("video"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/1"),
						},
					},
				},
				{
					Name: []byte("video"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/2"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with redundant attributes",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<video src="http://example.com/1" width="320" height="240"></video>
						</li>
						<li>
							<video src="http://example.com/2" width="640" height="480"></video>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"video": {"src": {}}},
				options: nil,
			},
			wantTags: []Tag{
				{
					Name: []byte("video"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/1"),
						},
					},
				},
				{
					Name: []byte("video"),
					Attributes: []Attribute{
						{
							Name:  []byte("src"),
							Value: []byte("http://example.com/2"),
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "error",
			args: args{
				reader: iotest.TimeoutReader(strings.NewReader(`
					<ul>
						<li><a href="http://example.com/1">1</a></li>
						<li><a href="http://example.com/2">2</a></li>
					</ul>
				`)),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				options: nil,
			},
			wantTags: nil,
			wantErr:  assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			gotTags, gotErr :=
				SelectTags(data.args.reader, data.args.filters, data.args.options...)

			assert.Equal(test, data.wantTags, gotTags)
			data.wantErr(test, gotErr)
		})
	}
}
