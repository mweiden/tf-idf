package main

import (
	"fmt"
	"log"
	"os"
)

var (
	options = map[Option]string{
		Option{"-h", "--help"}: "show this message",
	}
)

type Option struct {
	shortStr string
	longStr  string
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [file1.txt] [file2.txt] ...\n", os.Args[0])
	for k, v := range options {
		fmt.Fprintf(os.Stderr, "\t%s, %s\t%s", k.shortStr, k.longStr, v)
	}
	fmt.Fprint(os.Stderr, "\n")
}

func validateArguments() []string {
	num_args := len(os.Args)
	if num_args < 2 {
		printUsage()
		os.Exit(1)
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage()
		os.Exit(0)
	}
	log.Printf("Files selected: %v", os.Args[1:])
	return os.Args[1:]
}

type OutputFile struct {
	path string
	file *os.File
}

func main() {
	log.SetOutput(os.Stdout)
	filePaths := validateArguments()

	var cp CorpusProcessor
	cp.Init(filePaths)
	cp.Process()
}
