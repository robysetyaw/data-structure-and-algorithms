# Struktur Data dan Algoritma Golang

Proyek ini berisi implementasi berbagai struktur data dan algoritma umum menggunakan Golang serta implementasi unit test

## Menjalankan Unit Test

Proyek ini menggunakan paket bawaan `testing` dari Go.

Untuk menjalankan **semua unit test dalam proyek**, gunakan perintah berikut:

> go test ./...

## Menjalankan Test pada Paket Tertentu

Sebagai contoh, untuk menjalankan test hanya pada paket stack:

> go test ./stack

Atau untuk menjalankan test pada paket twosums:

> go test ./twosums

Semua unit test mengikuti pola penulisan func(input, want, description) Pola ini memudahkan dalam mendefinisikan dan memahami setiap test case.

# -----

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
