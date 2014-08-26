package txd

import (
	"errors"
	"speter.net/go/exp/math/dec/inf"
)

var (
	NoEnoughBalance  = errors.New("no enough balance")
	AccountNotExists = errors.New("account not exists")
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

func (balances Balances) SendTo(from int64, to int64, amount Amount) error {
	// send money to another account
	var ok bool
	if _, ok = balances[from]; !ok {
		return AccountNotExists
	}
	if _, ok = balances[to]; !ok {
		balances[to] = NewBalance()
	}

	if balances[from].Get(amount.Currency).Cmp(amount) < 0 {
		return NoEnoughBalance
	}
	return nil
}
