# tf-idf
![Go](https://github.com/mweiden/tf-idf/workflows/Go/badge.svg)

`tf-idf` calculates the [TF-IDF](https://en.wikipedia.org/wiki/Tf%E2%80%93idf) score of each word in a document given a corpus of documents.

The input to `tf-idf` is a list of text documents. The output is a set of files, named the same as the input document with the postfix `.tfidf`. Each output file is a TSV, listing each word in the corpus with its TF-IDF score.

## Building

To build `tf-idf`, simply run the following command.

```bash
$ make
```

### Testing

To run the project's unit tests, run the following command.

```bash
$ make test
```

## Use

To see usage instructions for `tf-idf` use the `-h` option, use `--help` option, or simply provide no input as shown below.

```bash
$ ./tf-idf --help
Usage ./tf-idf:
        ./tf-idf [file1.txt] [file2.txt] ...
        -h, --help      show this message
```
