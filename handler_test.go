package lab2

import (
	"bytes"
	. "gopkg.in/check.v1"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) { TestingT(t) }

type MySuiteHand struct{}

var _ = Suite(&MySuiteHand{})

func (s *MySuiteHand) TestComputeHandler_Compute(c *C) {
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  strings.NewReader("A B C * + D +"),
		Output: &output,
	}
	err := handler.Compute()
	if err != nil {
		c.Errorf(magicError)
	}
	str := output.String()
	c.Check(str, Equals, "++A*BCD")
}

func (s *MySuiteHand) TestComputeHandler_ComputeWrongInput(c *C) {
	var output bytes.Buffer

	handler := ComputeHandler{
		Input:  strings.NewReader("SOMETHING BAD"),
		Output: &output,
	}
	err := handler.Compute()

	c.Check(err, NotNil)
}
