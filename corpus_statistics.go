package main

import (
	"io"
	"math"
)

type WordKey struct {
	word     string
	document string
}

type CorpusStatistics struct {
	documentCount map[string]int32
	termFrequency map[WordKey]float64
	documents     []string
}

func (cs *CorpusStatistics) Init() {
	cs.documentCount = make(map[string]int32)
	cs.termFrequency = make(map[WordKey]float64)
}

func (cs *CorpusStatistics) AddDocument(scan TokenScanner) {
	cs.documents = append(cs.documents, scan.document)
	token, err := scan.Next()
	documentWords := make(map[string]bool)
	documentKeys := make(map[WordKey]bool)

	// term frequency numerator
	numTerms := 0.0
	for err != io.EOF {
		key := WordKey{token, scan.document}
		_, ok := cs.termFrequency[key]
		if ok {
			cs.termFrequency[key] += 1
		} else {
			cs.termFrequency[key] = 1
		}
		numTerms += 1.0

		documentWords[token] = true
		documentKeys[key] = true

		token, err = scan.Next()
	}
	for key, _ := range documentKeys {
		cs.termFrequency[key] /= numTerms
	}

	// count documents containing each word
	for word, _ := range documentWords {
		_, ok := cs.documentCount[word]
		if ok {
			cs.documentCount[word] += 1
		} else {
			cs.documentCount[word] = 1
		}
	}
}

func (cs *CorpusStatistics) TfIdf() map[WordKey]float64 {
	tfidf := make(map[WordKey]float64)
	N := float64(len(cs.documents))

	// zero all initial values
	for word, _ := range cs.documentCount {
		for _, doc := range cs.documents {
			tfidf[WordKey{word, doc}] = 0.0
		}
	}
	// calculate tf-idf
	for wordKey, _ := range cs.termFrequency {
		tf := float64(cs.termFrequency[wordKey])
		idf := math.Log10(N / float64(cs.documentCount[wordKey.word]))
		tfidf[wordKey] = tf * idf
	}
	return tfidf
}
