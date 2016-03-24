package webqueue_test

import (
	. "github.com/Enrise/webqueue"
	check "gopkg.in/check.v1"
	"io/ioutil"
	"net/http"
	"strings"
)

type UtilSuite struct{}

var _ = check.Suite(&UtilSuite{})

func (s *UtilSuite) TestRogueRead(c *check.C) {
	request, _ := http.NewRequest("POST", "/", strings.NewReader("Foobar"))

	result := RogueRead(&request.Body)
	c.Assert(result, check.FitsTypeOf, string(""))
	c.Assert(result, check.Equals, "Foobar")

	readerContents, _ := ioutil.ReadAll(request.Body)
	c.Assert(string(readerContents), check.Equals, "Foobar")

	// Now it's really empty
	readerContents, _ = ioutil.ReadAll(request.Body)
	c.Assert(string(readerContents), check.Equals, "")
}
