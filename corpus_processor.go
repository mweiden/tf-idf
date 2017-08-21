package main

import (
	"fmt"
	"log"
	"os"
)

type CorpusProcessor struct {
	inputFilepaths []string
	corpusStats    CorpusStatistics
}

func (cp *CorpusProcessor) Init(inputFilepaths []string) *CorpusProcessor {
	cp.inputFilepaths = inputFilepaths
	var cs CorpusStatistics
	cs.Init()
	cp.corpusStats = cs
	return cp
}

func (cp *CorpusProcessor) Process() {
	for _, filePath := range cp.inputFilepaths {
		var ts TokenScanner
		ts.FromFile(filePath)
		cp.corpusStats.AddDocument(ts)
		ts.Close()
	}

	tfidfMap := cp.corpusStats.TfIdf()

	outputFiles := make(map[string]OutputFile)
	for _, filePath := range cp.inputFilepaths {
		outputPath := filePath + ".tfidf"
		f, err := os.Create(outputPath)
		check(err)
		outputFiles[filePath] = OutputFile{outputPath, f}
	}

	for wordKey, tfidf := range tfidfMap {
		line := []byte(fmt.Sprintf("%s\t%f\n", wordKey.word, tfidf))
		outputFiles[wordKey.document].file.Write(line)
	}

	for originFilepath, f := range outputFiles {
		log.Printf("Wrote output for %s -> %s", originFilepath, f.path)
		f.file.Close()
	}
}
