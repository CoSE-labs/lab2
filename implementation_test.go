package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPostfixToPrefix(c *C) {
	inputs := []string{
		"A B C * + D +",
		"A B + C D + *",
		"A B * C D * +",
		"A B + C + D +",
		"6 9 + 4 2 * 4 2 ^ + +",
		"a t + b a c + * c d + ^ *",
		"A B C * + D +@",
		"text text",
		"",
	}
	expected := []string{
		"++A*BCD",
		"*+AB+CD",
		"+*AB*CD",
		"+++ABCD",
		"++69+*42^42",
		"*+at^*b+ac+cd",
		"Could not convert.\n",
		"Could not convert.\n",
		"Could not convert.\n",
	}

	for ind := range inputs {
		res, err := postfixToPrefix(inputs[ind])
		if err == nil {
			c.Check(res, Equals, expected[ind])
		} else {
			c.Check(err, ErrorMatches, expected[ind])
		}
	}
}

func Example_postfixToPrefix() {
	res, err := postfixToPrefix("4 8 3 * +")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	// Output:
	// +4*83
}
