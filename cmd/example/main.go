package main

import (
	"flag"
	"io"
	"os"
	"strings"

	lab2 "github.com/CoSE-labs/lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Expression from file")
	outputFile      = flag.String("o", "", "Filename to print result")
)

var mainComputeHandler lab2.ComputeHandler

func main() {

	flag.Parse()

	var input io.Reader = nil
	var output io.Writer = nil

	if *inputExpression != "" && *inputFile != "" {
		panic("Wrong format")
	}

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		fil, err := os.OpenFile(*inputFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		defer fil.Close()
		input = fil
	}

	if *outputFile != "" {
		fil, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		defer fil.Close()
		output = fil
	} else {
		output = os.Stdout
	}

	mainComputeHandler = lab2.ComputeHandler{Input: input, Output: output}

	err := mainComputeHandler.Compute()
	if err != nil {
		panic(err)
	}
}
