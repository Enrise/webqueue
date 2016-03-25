package webqueue

import (
	"bytes"
	"io"
	"io/ioutil"
)

// RogueRead allows reading the string contents from a stream while
// still being able to read from the ReadCloser.
func RogueRead(r *io.ReadCloser) string {
	buffer, _ := ioutil.ReadAll(*r)
	*r = ioutil.NopCloser(bytes.NewBuffer(buffer))

	return string(buffer)
}
