package byteutils

import (
	"unsafe"
)

// Copy ...
func Copy(bytes []byte) []byte {
	bytesCopy := make([]byte, len(bytes))
	copy(bytesCopy, bytes)

	return bytesCopy
}

// String ...
//
// Conversion from a byte slice to a string without memory allocation.
//
// See for explanations:
//   - https://github.com/golang/go/issues/25484
//   - implementation of https://golang.org/pkg/strings/#Builder.String
//
func String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes)) // nolint: gosec
}
