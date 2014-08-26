package txd

import (
	"speter.net/go/exp/math/dec/inf"
)

type Currency [3]byte

func (currency Currency) String() string {
	return string(currency[:])
}

func NewCurrency(c string) Currency {
	return [3]byte{c[0], c[1], c[2]}
}

type Amount struct {
	Value    *inf.Dec
	Currency Currency
}

func NewAmount(currency string, unscaled int64, scale int32) Amount {
	return Amount{inf.NewDec(unscaled, inf.Scale(scale)), NewCurrency(currency)}
}

func (amount Amount) String() string {
	return amount.Value.String() + amount.Currency.String()
}

func (this Amount) Equal(another Amount) bool {
	return this.Currency == another.Currency && this.Value.Cmp(another.Value) == 0
}

func (amount Amount) Neg() {
	amount.Value.Neg(amount.Value)
}

func (this Amount) Cmp(another Amount) int {
	return this.Value.Cmp(another.Value)
}
