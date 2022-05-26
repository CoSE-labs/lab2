package main

import (
	"flag"
	lab2 "github.com/CoSE-labs/lab2"
	"os"
	"strings"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Expression from file")
	outputFile      = flag.String("o", "", "Filename to print result")
)

var mainComputeHandler lab2.ComputeHandler

func main() {

	flag.Parse()

	if *inputExpression != "" {
		mainComputeHandler = lab2.ComputeHandler{Input: strings.NewReader(*inputExpression)}
	} else if *inputFile != "" {
		fil, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		defer fil.Close()
		mainComputeHandler = lab2.ComputeHandler{Input: fil}
	}
	if *outputFile != "" {
		fil, err := os.Open(*outputFile)
		if err != nil {
			panic(err)
		}
		defer fil.Close()
		mainComputeHandler = lab2.ComputeHandler{Output: fil}
	} else {
		mainComputeHandler = lab2.ComputeHandler{Output: os.Stdout}
	}

	err := mainComputeHandler.Compute()
	if err != nil {
		panic(err)
	}
}
