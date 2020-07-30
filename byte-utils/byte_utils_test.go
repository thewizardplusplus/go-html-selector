package byteutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(test *testing.T) {
	type args struct {
		bytes []byte
	}

	for _, data := range []struct {
		name         string
		args         args
		action       func(bytes []byte)
		wantOriginal []byte
		wantCopy     []byte
	}{
		{
			name: "nil",
			args: args{
				bytes: nil,
			},
			action:       func(bytes []byte) {},
			wantOriginal: nil,
			wantCopy:     []byte{},
		},
		{
			name: "empty",
			args: args{
				bytes: []byte{},
			},
			action:       func(bytes []byte) {},
			wantOriginal: []byte{},
			wantCopy:     []byte{},
		},
		{
			name: "nonempty/without changes",
			args: args{
				bytes: []byte("test"),
			},
			action:       func(bytes []byte) {},
			wantOriginal: []byte("test"),
			wantCopy:     []byte("test"),
		},
		{
			name: "nonempty/with changes",
			args: args{
				bytes: []byte("test"),
			},
			action:       func(bytes []byte) { bytes[2] = 'x' },
			wantOriginal: []byte("text"),
			wantCopy:     []byte("test"),
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got := Copy(data.args.bytes)
			data.action(data.args.bytes)

			assert.Equal(test, data.wantOriginal, data.args.bytes)
			assert.Equal(test, data.wantCopy, got)
		})
	}
}

func TestString(test *testing.T) {
	type args struct {
		bytes []byte
	}

	for _, data := range []struct {
		name         string
		args         args
		action       func(bytes []byte)
		wantOriginal []byte
		wantString   string
	}{
		{
			name: "nil",
			args: args{
				bytes: nil,
			},
			action:       func(bytes []byte) {},
			wantOriginal: nil,
			wantString:   "",
		},
		{
			name: "empty",
			args: args{
				bytes: []byte{},
			},
			action:       func(bytes []byte) {},
			wantOriginal: []byte{},
			wantString:   "",
		},
		{
			name: "nonempty/without changes",
			args: args{
				bytes: []byte("test"),
			},
			action:       func(bytes []byte) {},
			wantOriginal: []byte("test"),
			wantString:   "test",
		},
		{
			name: "nonempty/with changes",
			args: args{
				bytes: []byte("test"),
			},
			action:       func(bytes []byte) { bytes[2] = 'x' },
			wantOriginal: []byte("text"),
			wantString:   "text",
		},
	} {
		test.Run(data.name, func(test *testing.T) {
			got := String(data.args.bytes)
			data.action(data.args.bytes)

			assert.Equal(test, data.wantOriginal, data.args.bytes)
			assert.Equal(test, data.wantString, got)
		})
	}
}
