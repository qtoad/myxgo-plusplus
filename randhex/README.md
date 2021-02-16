# randhex

[![Build Status](https://travis-ci.org/frankenbeanies/randhex.svg?branch=master)](https://travis-ci.org/frankenbeanies/randhex) 
[![Coverage Status](https://coveralls.io/repos/github/frankenbeanies/randhex/badge.svg?branch=master)](https://coveralls.io/github/frankenbeanies/randhex?branch=master) 
[![Go Report Card](https://goreportcard.com/badge/github.com/frankenbeanies/randhex)](https://goreportcard.com/report/github.com/frankenbeanies/randhex)
[![GoDoc](https://godoc.org/github.com/frankenbeanies/uuid4?status.svg)](https://godoc.org/github.com/frankenbeanies/randhex)

A library for generating random hexadecimal color codes

## License

[MIT](LICENSE)

## Installation

```
$ go get github.com/frankenbeanies/randhex
```

## Usage

```go
import "github.com/frankenbeanies/randhex
```

## Methods

### New()

Generates a new random hexadecimal color code (RandHex)

```go
hex := randhex.New()
```

### String()

Provides a string representation of the RandHex prepended with '#'

```go
hexStr := randhex.New().String()
```

### Bytes()

Provides the byte representation of the RandHex

```go
hexBytes := randHex.New().Bytes()
```

### ParseString()
Parses string into a RandHex

```go
hex, _ := randhex.ParseString("#aaa")
hex, _ := randhex.ParseString("#AAAAAA")
hex, _ := randhex.ParseString("aaa")
_, err := randhex.ParseString("a") //error, bad length
_, err := randhex.ParseString("%aaa") //error, bad symbol
_, err := randhex.ParseString("#gaa") //error, not hex
```