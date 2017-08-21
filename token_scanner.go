package main

import (
	"bufio"
	"io"
	"os"
	"text/scanner"
	"unicode"
)

type TokenScanner struct {
	scanr    *scanner.Scanner
	file     *os.File
	document string
}

func (ts *TokenScanner) Close() error {
	return ts.file.Close()
}

func (ts *TokenScanner) FromFile(filePath string) *TokenScanner {
	ts.scanr, ts.file = scanFile(filePath)
	ts.document = filePath
	return ts
}

func (ts *TokenScanner) Next() (string, error) {
	nextString := ""
	var token rune = ' '
	for !unicode.IsLetter(token) && token != scanner.EOF {
		token = ts.scanr.Next()
	}
	for unicode.IsLetter(token) && token != scanner.EOF {
		nextString = nextString + string(unicode.ToLower(token))
		token = ts.scanr.Next()
	}
	if token == scanner.EOF {
		return nextString, io.EOF
	} else {
		return nextString, nil
	}
}

func scanFile(filePath string) (*scanner.Scanner, *os.File) {
	f, err := os.Open(filePath)
	check(err)
	reader := bufio.NewReader(f)
	scan := scanner.Scanner{}
	return scan.Init(reader), f
}
