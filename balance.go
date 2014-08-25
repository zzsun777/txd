package txd

import (
	"speter.net/go/exp/math/dec/inf"
)

type Balance map[Currency]Amount

func NewBalance() Balance {
	return make(Balance)
}

func (balance Balance) Get(currency Currency) Amount {
	if amount, ok := balance[currency]; ok {
		return amount
	}
	return Amount{Currency: currency, Value: inf.NewDec(0, 0)}
}

func (balance Balance) Inc(amount Amount) Amount {
	currency := amount.Currency

	if _, ok := balance[currency]; ok {
		balance[currency].Value.Add(balance[currency].Value, amount.Value)
	} else {
		balance[currency] = amount
	}
	return balance[currency]
}

type Balances map[int64]Balance

func (blances Balances) SendTo(from int64, to int64) {

}
