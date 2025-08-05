# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is the faker library for Go - a data generation library for creating fake/mock data for testing, database seeding, and anonymization. It's heavily inspired by PHP's Faker.

The library requires Go >= 1.22 (as per go.mod) and uses Go's math/rand/v2 package for random generation.

## Core Architecture

The library follows a modular design where each data type has its own file and corresponding test file:
- `faker.go` - Main Faker struct and GeneratorInterface that provides thread-safe random generation
- Each data category (e.g., `person.go`, `address.go`, `internet.go`) implements its own struct with methods
- All generators are accessed through the main Faker instance created with `faker.New()`
- The library uses struct tags (`fake:"..."`) to support generating fake data directly into structs

Key pattern: Each generator returns a new instance that shares the same underlying random generator for consistency and thread safety.

## Common Commands

### Testing
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with benchmarks
go test -v -bench=. ./...

# Run a specific test
go test -run TestPersonName -v

# Run tests multiple times (as done in CI)
for i in 1 2 3; do go test -v -bench=. ./... && break || if [ $i -eq 3 ]; then exit 1; fi; done
```

### Code Quality
```bash
# Format code
gofmt -w .

# Check for issues
go vet ./...

# Build the library
go build ./...
```

Note: There are known struct tag syntax issues in `struct_test.go` (lines 239, 248, 274) that cause `go vet` warnings. These are intentional test cases for malformed tags.

## Development Workflow

1. The library is purely functional with no external dependencies beyond Go standard library
2. Each data generator should have comprehensive tests in its corresponding `_test.go` file
3. Thread safety is handled at the Faker struct level through the threadSafeRand wrapper
4. When adding new generators, follow the existing pattern of creating a new struct type with methods
5. The main branch is `master` (not `main`)