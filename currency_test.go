package currency_test

import (
	"encoding"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/talav/currency"
)

// iso builds an ISO 4217 currency that the test knows to exist.
func iso(t *testing.T, code string) currency.Currency {
	t.Helper()

	c, err := currency.NewCurrencyFromString(code)
	require.NoError(t, err)

	return c
}

// custom builds a currency outside ISO 4217 that the test knows to be valid.
func custom(t *testing.T, code string, decimals int) currency.Currency {
	t.Helper()

	c, err := currency.NewCurrency(code, decimals)
	require.NoError(t, err)

	return c
}

func TestNewCurrency(t *testing.T) {
	t.Parallel()

	const codeErr = "must be 2 to 12 upper-case letters or digits"

	tests := []struct {
		name     string
		code     string
		decimals int
		wantCode string
		wantErr  string
	}{
		{name: "bitcoin", code: "BTC", decimals: 8, wantCode: "BTC"},
		{name: "four-letter code", code: "USDT", decimals: 6, wantCode: "USDT"},
		{name: "most decimals", code: "ETH", decimals: 18, wantCode: "ETH"},
		{name: "no decimals", code: "PTS", decimals: 0, wantCode: "PTS"},
		{name: "code starting with a digit", code: "1INCH", decimals: 18, wantCode: "1INCH"},
		{name: "shortest code", code: "X1", decimals: 2, wantCode: "X1"},
		{name: "longest code", code: "ABCDEFGHIJKL", decimals: 2, wantCode: "ABCDEFGHIJKL"},
		{name: "withdrawn national currency", code: "DEM", decimals: 2, wantCode: "DEM"},
		{name: "lower case", code: "btc", decimals: 8, wantCode: "BTC"},
		{name: "mixed case", code: "Btc", decimals: 8, wantCode: "BTC"},
		{name: "surrounding white space", code: " usdt\n", decimals: 6, wantCode: "USDT"},
		{name: "empty", code: "", decimals: 2, wantErr: `code "" ` + codeErr},
		{name: "one letter", code: "X", decimals: 2, wantErr: `code "X" ` + codeErr},
		{name: "thirteen letters", code: "ABCDEFGHIJKLM", decimals: 2, wantErr: `code "ABCDEFGHIJKLM" ` + codeErr},
		{name: "only white space", code: "  ", decimals: 2, wantErr: `code "" ` + codeErr},
		{name: "space inside", code: "BT C", decimals: 8, wantErr: `code "BT C" ` + codeErr},
		{name: "hyphen", code: "BTC-USD", decimals: 8, wantErr: `code "BTC-USD" ` + codeErr},
		{name: "non-ASCII letter", code: "BTÇ", decimals: 8, wantErr: `code "BTÇ" ` + codeErr},
		{name: "non-ASCII lower-case letter", code: "btç", decimals: 8, wantErr: `code "BTç" ` + codeErr},
		{name: "negative decimals", code: "BTC", decimals: -1, wantErr: "decimals -1 must be in 0..18"},
		{name: "too many decimals", code: "BTC", decimals: 19, wantErr: "decimals 19 must be in 0..18"},
		{name: "ISO code in lower case", code: "usd", decimals: 2, wantErr: `code "USD" is an ISO 4217 currency: use NewCurrencyFromString`},
		{name: "ISO code", code: "USD", decimals: 2, wantErr: `code "USD" is an ISO 4217 currency: use NewCurrencyFromString`},
		{name: "ISO code with other decimals", code: "JPY", decimals: 2, wantErr: `code "JPY" is an ISO 4217 currency: use NewCurrencyFromString`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := currency.NewCurrency(tt.code, tt.decimals)
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				assert.Equal(t, currency.Currency{}, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantCode, got.Code())
			assert.Equal(t, tt.wantCode, got.String())
			assert.Equal(t, tt.decimals, got.Decimals())
			assert.Zero(t, got.NumericCode())
		})
	}
}

func TestNewCurrencyFromString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		code         string
		wantCode     string
		wantNumeric  int
		wantDecimals int
		wantErr      bool
	}{
		{name: "lower case", code: "usd", wantCode: "USD", wantNumeric: 840, wantDecimals: 2},
		{name: "mixed case", code: "Eur", wantCode: "EUR", wantNumeric: 978, wantDecimals: 2},
		{name: "surrounding white space", code: " JPY\t", wantCode: "JPY", wantNumeric: 392, wantDecimals: 0},
		{name: "padded lower case", code: "gbp  ", wantCode: "GBP", wantNumeric: 826, wantDecimals: 2},
		{name: "two decimals", code: "USD", wantCode: "USD", wantNumeric: 840, wantDecimals: 2},
		{name: "euro", code: "EUR", wantCode: "EUR", wantNumeric: 978, wantDecimals: 2},
		{name: "no decimals", code: "JPY", wantCode: "JPY", wantNumeric: 392, wantDecimals: 0},
		{name: "three decimals", code: "BHD", wantCode: "BHD", wantNumeric: 48, wantDecimals: 3},
		{name: "four decimals", code: "CLF", wantCode: "CLF", wantNumeric: 990, wantDecimals: 4},
		{name: "numeric code below 100", code: "ALL", wantCode: "ALL", wantNumeric: 8, wantDecimals: 2},
		{name: "no minor unit", code: "XAU", wantCode: "XAU", wantNumeric: 959, wantDecimals: 0},
		{name: "fund", code: "USN", wantCode: "USN", wantNumeric: 997, wantDecimals: 2},
		{name: "first in the table", code: "AED", wantCode: "AED", wantNumeric: 784, wantDecimals: 2},
		{name: "last in the table", code: "ZWG", wantCode: "ZWG", wantNumeric: 924, wantDecimals: 2},
		{name: "empty", code: "", wantErr: true},
		{name: "two letters", code: "US", wantErr: true},
		{name: "space inside", code: "US D", wantErr: true},
		{name: "letter that upper-cases to an ASCII one", code: "uſd", wantErr: true},
		{name: "numeric code as text", code: "840", wantErr: true},
		{name: "cryptocurrency", code: "BTC", wantErr: true},
		{name: "withdrawn currency", code: "DEM", wantErr: true},
		{name: "before the first", code: "AAA", wantErr: true},
		{name: "after the last", code: "ZZZ", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := currency.NewCurrencyFromString(tt.code)
			if tt.wantErr {
				require.EqualError(t, err, `invalid currency "`+tt.code+`": not an ISO 4217 code`)
				assert.Equal(t, currency.Currency{}, got)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.wantCode, got.Code())
			assert.Equal(t, tt.wantNumeric, got.NumericCode())
			assert.Equal(t, tt.wantDecimals, got.Decimals())
		})
	}
}

func TestCurrencyEquality(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		left  currency.Currency
		right currency.Currency
		want  bool
	}{
		{name: "same ISO currency", left: iso(t, "USD"), right: iso(t, "USD"), want: true},
		{name: "different ISO currencies", left: iso(t, "USD"), right: iso(t, "CAD"), want: false},
		{name: "same custom currency", left: custom(t, "BTC", 8), right: custom(t, "BTC", 8), want: true},
		{name: "different custom currencies", left: custom(t, "BTC", 8), right: custom(t, "USDT", 6), want: false},
		{name: "same code, different decimals", left: custom(t, "BTC", 8), right: custom(t, "BTC", 6), want: false},
		{name: "ISO against the zero value", left: iso(t, "USD"), right: currency.Currency{}, want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.left == tt.right)
		})
	}
}

func TestCurrencyText(t *testing.T) {
	t.Parallel()

	var (
		_ encoding.TextMarshaler   = currency.Currency{}
		_ encoding.TextUnmarshaler = &currency.Currency{}
	)

	tests := []struct {
		name      string
		currency  currency.Currency
		want      string
		readsBack bool
	}{
		{name: "ISO currency", currency: iso(t, "USD"), want: "USD", readsBack: true},
		{name: "ISO currency with no decimals", currency: iso(t, "JPY"), want: "JPY", readsBack: true},
		{name: "custom currency", currency: custom(t, "BTC", 8), want: "BTC", readsBack: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, tt.currency.String())

			text, err := tt.currency.MarshalText()
			require.NoError(t, err)
			assert.Equal(t, tt.want, string(text))

			value, err := tt.currency.Value()
			require.NoError(t, err)
			assert.Equal(t, tt.want, value)

			var parsed currency.Currency

			err = parsed.UnmarshalText(text)
			if !tt.readsBack {
				require.Error(t, err)
				assert.Equal(t, currency.Currency{}, parsed)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.currency, parsed)
		})
	}
}

func TestCurrencyJSON(t *testing.T) {
	t.Parallel()

	type payload struct {
		Quote    currency.Currency            `json:"quote"`
		Balances map[currency.Currency]string `json:"balances"`
	}

	in := payload{
		Quote:    iso(t, "USD"),
		Balances: map[currency.Currency]string{iso(t, "EUR"): "12.50"},
	}

	data, err := json.Marshal(in)
	require.NoError(t, err)
	assert.JSONEq(t, `{"quote":"USD","balances":{"EUR":"12.50"}}`, string(data))

	var out payload

	require.NoError(t, json.Unmarshal(data, &out))
	assert.Equal(t, in, out)
}

func TestCurrencyScan(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		src     any
		want    currency.Currency
		wantErr string
	}{
		{name: "string", src: "EUR", want: iso(t, "EUR")},
		{name: "bytes", src: []byte("JPY"), want: iso(t, "JPY")},
		{name: "null", src: nil, wantErr: "invalid currency: NULL"},
		{name: "unsupported type", src: int64(840), wantErr: "invalid currency: unsupported type int64"},
		{name: "invalid string", src: "EURO", wantErr: `invalid currency "EURO": not an ISO 4217 code`},
		{name: "bytes in lower case", src: []byte("eur"), want: iso(t, "EUR")},
		{name: "padded string", src: "CHF  ", want: iso(t, "CHF")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			before := iso(t, "USD")
			got := before

			err := got.Scan(tt.src)
			if tt.wantErr != "" {
				require.EqualError(t, err, tt.wantErr)
				assert.Equal(t, before, got, "a failed scan must leave the currency unchanged")

				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
