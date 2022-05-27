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
	tests := map[string]string{
		"A B C * + D +":             "++A*BCD",
		"A B + C D + *":             "*+AB+CD",
		"A B * C D * +":             "+*AB*CD",
		"A B + C + D +":             "+++ABCD",
		"6 9 + 4 2 * 4 2 ^ + +":     "++69+*42^42",
		"a t + b a c + * c d + ^ *": "*+at^*b+ac+cd",
		"A B C * + D +@":            "Could not convert.\n",
		"text text":                 "Could not convert.\n",
		"":                          "Could not convert.\n",
	}

	for actual, expected := range tests {
		res, err := postfixToPrefix(actual)
		if err == nil {
			c.Check(res, Equals, expected)
		} else {
			c.Check(err, ErrorMatches, expected)
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
