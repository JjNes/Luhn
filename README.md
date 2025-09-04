# Luhn Validator Package
[![Go](https://github.com/jjnes/luhn/actions/workflows/go.yml/badge.svg)](https://github.com/jjnes/luhn/actions)
[![GitHub tag](https://img.shields.io/github/v/tag/jjnes/luhn?label=version&sort=semver)](https://github.com/jjnes/luhn/tags)
[![codecov](https://codecov.io/gh/JjNes/Luhn/graph/badge.svg?token=S4Q4YYVR0B)](https://codecov.io/gh/JjNes/Luhn)

This package provides functionality to validate numbers using the Luhn algorithm, which is commonly used to validate credit card numbers and other identification numbers.

## Install

```bash
go get github.com/jjnes/luhn
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/jjnes/luhn"
)

func main() {
	// Example usage for integer
	number := 1230
	isValid := luhn.IsValidInt(number)
	fmt.Println("Is valid number:", isValid)

	// Example usage for string
	numberStr := "1230"
	isValidStr := luhn.IsValidStr(numberStr)
	fmt.Println("Is valid string:", isValidStr)

	// Example usage for generate number
	prefix := 123
	fullNumber := luhn.GenerateLuhnNumber(prefix)
	isValidFullNumber := luhn.IsValidInt(fullNumber)
	fmt.Printf("Is valid %d: %t", fullNumber, isValidFullNumber)
}
```

## Testing

Tests are provided for the Luhn validation functions using the Go testing framework. To run the tests, use:

```sh
cd luhn
go test
```

The tests cover valid and invalid numbers, including checks for non-numeric strings.