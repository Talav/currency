# currency

[![tag](https://img.shields.io/github/tag/talav/currency.svg)](https://github.com/talav/currency/tags)
[![Go Reference](https://pkg.go.dev/badge/github.com/talav/currency.svg)](https://pkg.go.dev/github.com/talav/currency)
[![Go Report Card](https://goreportcard.com/badge/github.com/talav/currency)](https://goreportcard.com/report/github.com/talav/currency)
[![CI](https://github.com/talav/currency/workflows/Currency%20CI/badge.svg)](https://github.com/talav/currency/actions)
[![codecov](https://codecov.io/gh/talav/currency/graph/badge.svg)](https://codecov.io/gh/talav/currency)
[![License](https://img.shields.io/github/license/talav/currency)](./LICENSE)

A currency value for Go: a code, its ISO 4217 number and its decimal places, that cannot hold an invalid value.

## Features

- **ISO 4217 built in**: all current currencies and funds, with numeric codes and decimal places
- **Custom currencies**: cryptocurrencies, withdrawn currencies or anything else, as plain values
- **Valid by construction**: constructors return an error instead of a malformed code
- **One code, one meaning**: an ISO code always carries the ISO data and cannot be redefined
- **No global registry**: a currency is a value, with no registration step and no shared state
- **Comparable**: values work with `==` and as map keys
- **Text, JSON and SQL**: `encoding.TextMarshaler`, `driver.Valuer` and `sql.Scanner`
- **No dependencies** (except testing)

## Installation

```bash
go get github.com/talav/currency
```

Requires Go 1.27 or later.

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/talav/currency"
)

func main() {
    usd, err := currency.NewCurrencyFromString("USD")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(usd.Code())        // USD
    fmt.Println(usd.NumericCode()) // 840
    fmt.Println(usd.Decimals())    // 2

    btc, err := currency.NewCurrency("BTC", 8)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(btc, btc.Decimals()) // BTC 8
}
```

## Usage

### ISO 4217 currencies

```go
jpy, err := currency.NewCurrencyFromString("JPY")

jpy.Code()        // "JPY"
jpy.NumericCode() // 392
jpy.Decimals()    // 0
```

The code may be in any letter case, and white space around it is ignored, so `"usd"` and `" USD "` both give USD. It must be on list one of ISO 4217, the current currencies and funds. Anything else is an error:

```go
_, err := currency.NewCurrencyFromString("BTC")
// invalid currency "BTC": not an ISO 4217 code
```

ISO 4217 gives no smallest unit for a few codes, such as `XAU` for gold. `Decimals` returns 0 for them.

### Other currencies

A currency that ISO 4217 does not list is built from its code and its number of decimal places:

```go
btc, err := currency.NewCurrency("BTC", 8)
usdt, err := currency.NewCurrency("USDT", 6)
dem, err := currency.NewCurrency("DEM", 2) // withdrawn in 2002
```

The code is 2 to 12 letters or digits, stored in upper case, and the decimal places are 0 to 18. Such a currency has no numeric code: `NumericCode` returns 0.

An ISO code cannot be redefined, so that `USD` never means two different things:

```go
_, err := currency.NewCurrency("USD", 4)
// code "USD" is an ISO 4217 currency: use NewCurrencyFromString
```

### Comparing

Two currencies are equal when their code and decimal places are equal:

```go
a, _ := currency.NewCurrencyFromString("USD")
b, _ := currency.NewCurrencyFromString("USD")

a == b // true
```

### JSON and SQL

A currency is written as its code and works as a map key:

```go
type Account struct {
    Base currency.Currency `json:"base"`
}

// {"base":"USD"}
```

Reading a code back gives an ISO 4217 currency. A code alone does not say how many decimal places a currency outside ISO 4217 has, so decoding `"BTC"` is an error. An application that stores custom currencies reads the code as a string and looks the currency up in its own list.

`Scan` accepts a `string` or a `[]byte`. `NULL` is an error; scan a nullable column into a `sql.Null[currency.Currency]`.

## Zero Value

A currency is built by a constructor, but Go still lets anyone write `currency.Currency{}`. That value is not a currency. It stands for a missing one: `IsValid` returns false, and other methods are meaningless on it.

## ISO 4217 Data

The table in [`iso4217.go`](./iso4217.go) follows list one of ISO 4217 as published by the standard's maintenance agency on 2026-09-17. Withdrawn currencies are not included.

## Testing

```bash
go test -race -cover ./...
```

## License

MIT License. See [LICENSE](./LICENSE).
