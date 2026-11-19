// Package currency provides Currency: a currency identified by its code,
// together with the number of decimal places its amounts are kept to.
//
// The currencies of ISO 4217 are built in. Each is a named value, such as USD,
// and comes from NewCurrencyFromString. Any other currency, such as a
// cryptocurrency or a withdrawn national currency, comes from NewCurrency.
package currency

import (
	"cmp"
	"database/sql/driver"
	"errors"
	"fmt"
	"slices"
	"strings"
)

const (
	// A code is short enough to be a ticker and long enough for the longest
	// codes in use outside ISO 4217.
	minCodeLen = 2
	maxCodeLen = 12

	// maxDecimals is the largest number of decimal places a currency may
	// have. It fits the finest unit in common use, 18 places.
	maxDecimals = 18
)

// A Currency is a unit of money: a code, the ISO 4217 number if it has one,
// and the number of decimal places of its smallest unit.
//
// A Currency is a named value or comes from a constructor, so its code is well
// formed and a code that ISO 4217 lists always carries the number and decimal
// places that ISO 4217 gives it. Currencies are comparable with == and usable
// as map keys.
//
// The zero value is not a currency. It stands for a missing one, IsValid
// reports false for it, and other methods are meaningless on it.
type Currency struct {
	// code identifies the currency: three upper-case letters for an ISO 4217
	// currency, 2 to 12 upper-case letters or digits for any other.
	code string

	// numericCode is the three-digit ISO 4217 number, 1 to 999. It is 0 for a
	// currency that ISO 4217 does not list.
	numericCode int

	// decimals is the number of decimal places of the smallest unit of the
	// currency, 0 to 18: 2 for a currency divided into hundredths.
	decimals int
}

// NewCurrency returns a currency that ISO 4217 does not list, with the given
// code and number of decimal places. It has no numeric code.
//
// The code is read as NewCurrencyFromString reads one: in any letter case and
// with surrounding white space ignored. It returns an error if the code is
// not 2 to 12 letters or digits, if decimals is not in 0 to 18, or if
// ISO 4217 lists the code: such a currency comes from NewCurrencyFromString,
// so that one code never means two different currencies.
func NewCurrency(code string, decimals int) (Currency, error) {
	c := Currency{code: normalizeCode(code), decimals: decimals}
	if err := c.validate(); err != nil {
		return Currency{}, err
	}

	return c, nil
}

// NewCurrencyFromString returns the ISO 4217 currency with the code s, such
// as USD. The code may be in any letter case, and white space around it is
// ignored: "usd" and " USD " both give USD.
//
// It returns an error if ISO 4217 does not list the code among its current
// currencies and funds.
func NewCurrencyFromString(s string) (Currency, error) {
	c, ok := lookupISO(normalizeCode(s))
	if !ok {
		return Currency{}, fmt.Errorf("invalid currency %q: not an ISO 4217 code", s)
	}

	return c, nil
}

// IsValid reports whether c is a currency. It is false only for the zero
// value, which stands for a missing one; every Currency from a constructor
// is valid.
func (c Currency) IsValid() bool {
	return c != Currency{}
}

// Code returns the code of the currency, such as USD.
func (c Currency) Code() string { return c.code }

// NumericCode returns the ISO 4217 number of the currency, 1 to 999, which is
// written with three digits: 840 for USD, 008 for ALL. It returns 0 for a
// currency that ISO 4217 does not list.
func (c Currency) NumericCode() int { return c.numericCode }

// Decimals returns the number of decimal places of the smallest unit of the
// currency: 2 for USD, 0 for JPY, 3 for BHD.
//
// ISO 4217 gives no smallest unit for a few of its codes, such as XAU for
// gold. Decimals returns 0 for them.
func (c Currency) Decimals() int { return c.decimals }

// String returns the code of the currency.
func (c Currency) String() string { return c.code }

// MarshalText returns the code of the currency. It never fails. It makes
// Currency an encoding.TextMarshaler, which encoding/json uses for values and
// for map keys.
func (c Currency) MarshalText() ([]byte, error) {
	return []byte(c.code), nil
}

// UnmarshalText sets c to the ISO 4217 currency that NewCurrencyFromString
// reads from text. If the text is not such a code, it returns the error and
// leaves c as it was. It makes *Currency an encoding.TextUnmarshaler.
//
// A code alone does not say how many decimal places a currency outside
// ISO 4217 has, so the text of one cannot be read back.
func (c *Currency) UnmarshalText(text []byte) error {
	parsed, err := NewCurrencyFromString(string(text))
	if err != nil {
		return err
	}

	*c = parsed

	return nil
}

// Value returns the code of the currency for use as a query argument. It
// makes Currency a driver.Valuer.
func (c Currency) Value() (driver.Value, error) {
	return c.code, nil
}

// Scan sets c from a database column value. It makes *Currency a
// sql.Scanner.
//
// A string or []byte must be a code that NewCurrencyFromString accepts. NULL
// is an error; scan a nullable column into a sql.Null[Currency]. On error c
// is unchanged.
func (c *Currency) Scan(src any) error {
	switch value := src.(type) {
	case string:
		return c.UnmarshalText([]byte(value))
	case []byte:
		return c.UnmarshalText(value)
	case nil:
		return errors.New("invalid currency: NULL")
	default:
		return fmt.Errorf("invalid currency: unsupported type %T", src)
	}
}

func (c Currency) validate() error {
	if !isCode(c.code) {
		return fmt.Errorf("code %q must be %d to %d upper-case letters or digits", c.code, minCodeLen, maxCodeLen)
	}

	if c.decimals < 0 || c.decimals > maxDecimals {
		return fmt.Errorf("decimals %d must be in 0..%d", c.decimals, maxDecimals)
	}

	if _, ok := lookupISO(c.code); ok {
		return fmt.Errorf("code %q is an ISO 4217 currency: use NewCurrencyFromString", c.code)
	}

	return nil
}

// normalizeCode returns s as a code is stored: without surrounding white
// space and with its letters in upper case. Only ASCII letters change, so
// that no other letter turns into one a code may hold.
func normalizeCode(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 'a' + 'A'
		}

		return r
	}, strings.TrimSpace(s))
}

// isCode reports whether s has the form of a stored currency code.
func isCode(s string) bool {
	if len(s) < minCodeLen || len(s) > maxCodeLen {
		return false
	}

	for i := range len(s) {
		isLetter := s[i] >= 'A' && s[i] <= 'Z'
		isDigit := s[i] >= '0' && s[i] <= '9'

		if !isLetter && !isDigit {
			return false
		}
	}

	return true
}

// lookupISO returns the ISO 4217 currency with the given code. The table it
// searches is sorted by code.
func lookupISO(code string) (Currency, bool) {
	i, found := slices.BinarySearchFunc(isoCurrencies[:], code, func(c Currency, code string) int {
		return cmp.Compare(c.code, code)
	})
	if !found {
		return Currency{}, false
	}

	return isoCurrencies[i], true
}
