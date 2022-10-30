# Struct reflect example for Go

[![Build Status](https://cloud.drone.io/api/badges/danil/structreflectexample/status.svg)](https://cloud.drone.io/danil/structreflectexample)
[![Go Reference](https://pkg.go.dev/badge/github.com/danil/structreflectexample.svg)](https://pkg.go.dev/github.com/danil/structreflectexample)

Source files are distributed under the BSD-style license.

## About

The software is considered to be at a alpha level of readiness,
its extremely slow and allocates a lots of memory.

## Benchmark

```sh
go test -count=1 -race -bench=. -benchmem ./...
goos: linux
goarch: amd64
pkg: github.com/danil/structreflectexample
cpu: 11th Gen Intel(R) Core(TM) i7-1165G7 @ 2.80GHz
BenchmarkReflect/reflect-8          1000000000           0.0000169 ns/op           0 B/op          0 allocs/op
BenchmarkMarshal/marshal-8          1000000000           0.0000176 ns/op           0 B/op          0 allocs/op
PASS
ok      github.com/danil/structreflectexample   0.027s
```
