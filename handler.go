package lab2

import (
	"io"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	computable := make([]byte, 0)
	minBuff := make([]byte, 1)

	for {
		_, err := ch.Input.Read(minBuff)
		if err == io.EOF {
			break
		}

		computable = append(computable, minBuff[0])

	}

	res, err := postfixToPrefix(string(computable))
	if err != nil {
		return err
	}
	_, err = ch.Output.Write([]byte(res))
	if err != nil {
		return err
	}

	return nil
}
