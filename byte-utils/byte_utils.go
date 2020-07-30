package byteutils

import (
	"unsafe"
)

// Copy ...
func Copy(bytes []byte) []byte {
	if len(bytes) == 0 {
		return nil
	}

	bytesCopy := make([]byte, len(bytes))
	copy(bytesCopy, bytes)

	return bytesCopy
}

// String ...
//
// Conversion from a byte slice to a string without memory allocation.
//
// Attention! It doesn't produce a copy of bytes. It returns a mutable string
// (it'll change when original bytes will change).
//
// See for explanations:
//   - https://github.com/golang/go/issues/25484
//   - implementation of https://golang.org/pkg/strings/#Builder.String
//
func String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes)) // nolint: gosec
}
