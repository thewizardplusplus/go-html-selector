package htmlselector

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSelectTags(test *testing.T) {
	type args struct {
		reader  io.Reader
		filters OptimizedFilterGroup
		builder Builder
		options []Option
	}

	for _, data := range []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "success/with empty arguments",
			args: args{
				reader:  strings.NewReader(""),
				filters: nil,
				builder: new(MockBuilder),
				options: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with an empty reader",
			args: args{
				reader:  strings.NewReader(""),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				builder: new(MockBuilder),
				options: nil,
			},
			wantErr: assert.NoError,
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
				builder: new(MockBuilder),
				options: nil,
			},
			wantErr: assert.NoError,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("img")).Times(2)
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2")).
						Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.On("AddTag", []byte("img")).Times(2)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1.1")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2.1")).
						Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("video")).Times(2)

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("video")).Times(2)

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("video")).Times(2)

					return builder
				}(),
				options: nil,
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
				builder: new(MockBuilder),
				options: []Option{SkipEmptyTags()},
			},
			wantErr: assert.NoError,
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
				builder: new(MockBuilder),
				options: []Option{SkipEmptyTags()},
			},
			wantErr: assert.NoError,
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
				builder: new(MockBuilder),
				options: []Option{SkipEmptyTags()},
			},
			wantErr: assert.NoError,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("video")).Times(2)
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("poster"), []byte("http://example.com/1.1")).
						Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2")).
						Once()
					builder.
						On("AddAttribute", []byte("poster"), []byte("http://example.com/2.1")).
						Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("video")).Times(2)
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2")).
						Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("video")).Times(2)
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2")).
						Once()

					return builder
				}(),
				options: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with empty attributes/without skipping",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><a href="">1</a></li>
						<li><a href="">2</a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.On("AddAttribute", []byte("href"), []byte{}).Times(2)

					return builder
				}(),
				options: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with empty attributes/with skipping of empty attributes",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><a href="">1</a></li>
						<li><a href="">2</a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)

					return builder
				}(),
				options: []Option{SkipEmptyAttributes()},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with empty attributes" +
				"/with skipping of empty tags and attributes",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li><a href="">1</a></li>
						<li><a href="">2</a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				builder: new(MockBuilder),
				options: []Option{SkipEmptyTags(), SkipEmptyAttributes()},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with the universal tag/without skipping",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<a href="http://example.com/1" title="title #1">
								<img src="http://example.com/1.1" alt="alt #1" />
							</a>
						</li>
						<li>
							<a href="http://example.com/2" title="title #2">
								<img src="http://example.com/2.1" alt="alt #2" />
							</a>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{
					UniversalTag: {"title": {}, "alt": {}, "src": {}},
					"a":          {"href": {}},
				},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("ul")).Once()
					builder.On("AddTag", []byte("li")).Times(2)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.On("AddTag", []byte("img")).Times(2)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.On("AddAttribute", []byte("title"), []byte("title #1")).Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1.1")).
						Once()
					builder.On("AddAttribute", []byte("alt"), []byte("alt #1")).Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()
					builder.On("AddAttribute", []byte("title"), []byte("title #2")).Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2.1")).
						Once()
					builder.On("AddAttribute", []byte("alt"), []byte("alt #2")).Once()

					return builder
				}(),
				options: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with the universal tag/with skipping",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<a href="http://example.com/1" title="title #1">
								<img src="http://example.com/1.1" alt="alt #1" />
							</a>
						</li>
						<li>
							<a href="http://example.com/2" title="title #2">
								<img src="http://example.com/2.1" alt="alt #2" />
							</a>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{
					UniversalTag: {"title": {}, "alt": {}, "src": {}},
					"a":          {"href": {}},
				},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.On("AddTag", []byte("img")).Times(2)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.On("AddAttribute", []byte("title"), []byte("title #1")).Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/1.1")).
						Once()
					builder.On("AddAttribute", []byte("alt"), []byte("alt #1")).Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()
					builder.On("AddAttribute", []byte("title"), []byte("title #2")).Once()
					builder.
						On("AddAttribute", []byte("src"), []byte("http://example.com/2.1")).
						Once()
					builder.On("AddAttribute", []byte("alt"), []byte("alt #2")).Once()

					return builder
				}(),
				options: []Option{SkipEmptyTags()},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with a text builder/without skipping",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>link #1: <a href="http://example.com/1">one</a></li>
						<li>link #2: <a href="http://example.com/2">two</a></li>
						<li>link #3: <a href="http://example.com/3"></a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(3)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/3")).
						Once()

					textBuilder := new(MockTextBuilder)
					textBuilder.On("AddText", []byte("\n"+strings.Repeat("\t", 4))).Once()
					textBuilder.On("AddText", []byte("\n"+strings.Repeat("\t", 5))).Times(2)
					textBuilder.On("AddText", []byte("\n"+strings.Repeat("\t", 6))).Times(3)
					textBuilder.On("AddText", []byte("link #1: ")).Once()
					textBuilder.On("AddText", []byte("one")).Once()
					textBuilder.On("AddText", []byte("link #2: ")).Once()
					textBuilder.On("AddText", []byte("two")).Once()
					textBuilder.On("AddText", []byte("link #3: ")).Once()

					return MultiBuilder{builder, textBuilder}
				}(),
				options: nil,
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with a text builder/with skipping",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>link #1: <a href="http://example.com/1">one</a></li>
						<li>link #2: <a href="http://example.com/2">two</a></li>
						<li>link #3: <a href="http://example.com/3"></a></li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"href": {}}},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(3)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/3")).
						Once()

					textBuilder := new(MockTextBuilder)
					textBuilder.On("AddText", []byte("link #1: ")).Once()
					textBuilder.On("AddText", []byte("one")).Once()
					textBuilder.On("AddText", []byte("link #2: ")).Once()
					textBuilder.On("AddText", []byte("two")).Once()
					textBuilder.On("AddText", []byte("link #3: ")).Once()

					return MultiBuilder{builder, textBuilder}
				}(),
				options: []Option{SkipEmptyText()},
			},
			wantErr: assert.NoError,
		},
		{
			name: "success/with attributes overlapping between tags",
			args: args{
				reader: strings.NewReader(`
					<ul>
						<li>
							<a class="class #1" href="http://example.com/1" title="title #1">
								<img
									class="class #1.1"
									src="http://example.com/1.1"
									title="title #1.1" />
							</a>
						</li>
						<li>
							<a class="class #2" href="http://example.com/2" title="title #2">
								<img
									class="class #2.1"
									src="http://example.com/2.1"
									title="title #2.1" />
							</a>
						</li>
					</ul>
				`),
				filters: OptimizedFilterGroup{"a": {"class": {}}, "img": {"title": {}}},
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.On("AddTag", []byte("img")).Times(2)
					builder.On("AddAttribute", []byte("class"), []byte("class #1")).Once()
					builder.On("AddAttribute", []byte("class"), []byte("class #2")).Once()
					builder.On("AddAttribute", []byte("title"), []byte("title #1.1")).Once()
					builder.On("AddAttribute", []byte("title"), []byte("title #2.1")).Once()

					return builder
				}(),
				options: nil,
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
				builder: func() Builder {
					builder := new(MockBuilder)
					builder.On("AddTag", []byte("a")).Times(2)
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/1")).
						Once()
					builder.
						On("AddAttribute", []byte("href"), []byte("http://example.com/2")).
						Once()

					return builder
				}(),
				options: nil,
			},
			wantErr: assert.Error,
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			gotErr := SelectTags(
				data.args.reader,
				data.args.filters,
				data.args.builder,
				data.args.options...,
			)

			var builders []interface{}
			if multiBuilder, ok := data.args.builder.(MultiBuilder); ok {
				builders = []interface{}{multiBuilder.Builder, multiBuilder.TextBuilder}
			} else {
				builders = []interface{}{data.args.builder}
			}
			mock.AssertExpectationsForObjects(test, builders...)

			data.wantErr(test, gotErr)
		})
	}
}
