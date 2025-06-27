# Go Algorithms and Data Structures

This project contains implementations of common data structures and algorithms in Go (Golang).

## Running Unit Tests

This project uses the built-in testing package from Go.
You can run all tests in the project with:

> go test ./...

This will recursively discover and execute all *_test.go files.

## Running Tests in a Specific Package

For example, to run tests only in the stack package:

> go test ./stack

Or to run tests in the twosums package:

> go test ./twosums

All unit test had pattern **func(input, want, description)** This makes it easy to define and understand test cases.