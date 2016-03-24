package webqueue

import (
	"bytes"
	"io"
	"io/ioutil"
)

func RogueRead(r *io.ReadCloser) string {
	buffer, _ := ioutil.ReadAll(*r)
	*r = ioutil.NopCloser(bytes.NewBuffer(buffer))

	return string(buffer)
}
