# Ech0

[![GoDoc](https://img.shields.io/badge/api-Godoc-blue.svg)](https://pkg.go.dev/github.com/rickb777/ech0)
[![Build Status](https://travis-ci.org/rickb777/ech0.svg?branch=master)](https://travis-ci.org/rickb777/ech0/builds)
[![Code Coverage](https://img.shields.io/coveralls/rickb777/ech0.svg)](https://coveralls.io/r/rickb777/ech0)
[![Go Report Card](https://goreportcard.com/badge/github.com/rickb777/ech0)](https://goreportcard.com/report/github.com/rickb777/ech0)
[![Issues](https://img.shields.io/github/issues/rickb777/ech0.svg)](https://github.com/rickb777/ech0/issues)

Ech0 (pronounced "echo zero") is a logging adapter for `echo.Logger` that uses [github.com/rs/zerolog](https://github.com/rs/zerolog) as the logging backend instead of the default [github.com/labstack/gommon/log](https://github.com/labstack/gommon/tree/master/log)

# Why?

1. I like [Echo](https://echo.labstack.com/).
1. I like [zerolog](https://github.com/rs/zerolog).
1. I wanted to have *one* logging backend in my Echo apps.

# Installing

`go get -u github.com/rickb777/ech0`

# Benchmarks

```
goos: linux
goarch: amd64
pkg: github.com/rickb777/ech0
BenchmarkZeroFormat-8     	 2104641	       570 ns/op	      21 B/op	       2 allocs/op
BenchmarkZeroJSON-8       	 1260433	       911 ns/op	     520 B/op	       2 allocs/op
BenchmarkZero-8           	 2198152	       563 ns/op	      21 B/op	       2 allocs/op
BenchmarkGommonFormat-8   	  563232	      2078 ns/op	     464 B/op	      12 allocs/op
BenchmarkGommonJSON-8     	  517410	      2667 ns/op	     688 B/op	      16 allocs/op
BenchmarkGommon-8         	  558044	      2121 ns/op	     464 B/op	      12 allocs/op
PASS
```
