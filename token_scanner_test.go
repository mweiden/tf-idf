package main

import (
	"io"
	"testing"
)

func TestTokenScanner(t *testing.T) {
	t.Parallel()
	var ts TokenScanner
	defer ts.Close()
	ts.FromFile("test1.txt")

	expected := []string{"this", "is", "a", "a", "sample"}

	for _, w := range expected {
		word, err := ts.Next()
		if word != w {
			t.Errorf("Expected %s, got %v", word, w)
		}
		if err == io.EOF {
			t.Error("Should not have reached EOF")
		}
	}

	word, err := ts.Next()
	if word != "" {
		t.Error("Expected \"\", got", word)
	}
	if err != io.EOF {
		t.Error("Expected io.EOF, got", err)
	}
}
