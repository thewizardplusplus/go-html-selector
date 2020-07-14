package htmlselector

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectTags(test *testing.T) {
	type args struct {
		reader  io.Reader
		filters []Filter
	}

	for _, data := range []struct {
		name     string
		args     args
		wantTags []Tag
		wantErr  assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	} {
		test.Run(data.name, func(test *testing.T) {
			gotTags, gotErr := SelectTags(data.args.reader, data.args.filters)

			assert.Equal(test, data.wantTags, gotTags)
			data.wantErr(test, gotErr)
		})
	}
}
