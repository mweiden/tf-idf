package main

import (
	"math"
	"reflect"
	"testing"
)

func TestCorpusStatistics_Init(t *testing.T) {
	t.Parallel()
	var cs CorpusStatistics
	cs.Init()
	if len(cs.documents) != 0 {
		t.Error("documents should be empty")
	}
	if cs.documentCount == nil {
		t.Error("documentCount should not equal nil")
	}
	if cs.termFrequency == nil {
		t.Error("termFrequency should not equal nil")
	}
}

func TestCorpusStatistics_AddDocument(t *testing.T) {
	t.Parallel()
	var cs CorpusStatistics
	cs.Init()
	var ts TokenScanner
	defer ts.Close()
	ts.FromFile("test1.txt")
	cs.AddDocument(ts)
	ts.FromFile("test2.txt")
	cs.AddDocument(ts)

	expectedDocuments := []string{"test1.txt", "test2.txt"}
	if !reflect.DeepEqual(cs.documents, expectedDocuments) {
		t.Errorf("%v != %v", expectedDocuments, cs.documents)
	}

	expectedDocumentCount := map[string]int32{
		"this":    2,
		"is":      2,
		"a":       1,
		"another": 1,
		"example": 1,
		"sample":  1,
	}
	if !reflect.DeepEqual(cs.documentCount, expectedDocumentCount) {
		t.Errorf("%v != %v", expectedDocumentCount, cs.documentCount)
	}

	expectedTermFrequencies := map[WordKey]float64{
		WordKey{"this", "test1.txt"}:    1.0 / 5.0,
		WordKey{"is", "test1.txt"}:      1.0 / 5.0,
		WordKey{"a", "test1.txt"}:       2.0 / 5.0,
		WordKey{"sample", "test1.txt"}:  1.0 / 5.0,
		WordKey{"this", "test2.txt"}:    1.0 / 7.0,
		WordKey{"is", "test2.txt"}:      1.0 / 7.0,
		WordKey{"another", "test2.txt"}: 2.0 / 7.0,
		WordKey{"example", "test2.txt"}: 3.0 / 7.0,
	}
	if !reflect.DeepEqual(cs.termFrequency, expectedTermFrequencies) {
		t.Errorf("%v != %v", expectedTermFrequencies, cs.termFrequency)
	}
}

func TestCorpusStatistics_TfIdf(t *testing.T) {
	t.Parallel()
	var cs CorpusStatistics
	cs.Init()
	var ts TokenScanner
	defer ts.Close()
	ts.FromFile("test1.txt")
	cs.AddDocument(ts)
	ts.FromFile("test2.txt")
	cs.AddDocument(ts)

	tfidf := cs.TfIdf()

	expectedTfIdf := map[WordKey]float64{
		WordKey{"this", "test1.txt"}:    1.0 / 5.0 * math.Log10(2.0/2.0),
		WordKey{"is", "test1.txt"}:      1.0 / 5.0 * math.Log10(2.0/2.0),
		WordKey{"a", "test1.txt"}:       2.0 / 5.0 * math.Log10(2.0/1.0),
		WordKey{"another", "test1.txt"}: 0.0,
		WordKey{"sample", "test1.txt"}:  1.0 / 5.0 * math.Log10(2.0/1.0),
		WordKey{"example", "test1.txt"}: 0.0,
		WordKey{"this", "test2.txt"}:    1.0 / 7.0 * math.Log10(2.0/2.0),
		WordKey{"is", "test2.txt"}:      1.0 / 7.0 * math.Log10(2.0/2.0),
		WordKey{"a", "test2.txt"}:       0.0,
		WordKey{"another", "test2.txt"}: 2.0 / 7.0 * math.Log10(2.0/1.0),
		WordKey{"sample", "test2.txt"}:  0.0,
		WordKey{"example", "test2.txt"}: 3.0 / 7.0 * math.Log10(2.0/1.0),
	}
	if !reflect.DeepEqual(tfidf, expectedTfIdf) {
		t.Errorf("%v != %v", expectedTfIdf, tfidf)
	}
}
