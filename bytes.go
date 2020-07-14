package htmlselector

func copyBytes(bytes []byte) []byte {
	bytesCopy := make([]byte, len(bytes))
	copy(bytesCopy, bytes)

	return bytesCopy
}
