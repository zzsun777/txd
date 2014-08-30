package txd_test

import (
	"github.com/aisk/txd"
	"testing"
)

var (
	btc = txd.NewCurrency("BTC")
)

func TestNewBalance(t *testing.T) {
	b := new(txd.Balance)
	if !b.Get(btc).Equal(txd.NewAmount("BTC", 0, 0)) {
		t.Error()
	}
}

func TestIncBalance(t *testing.T) {
	b := txd.NewBalance()
	inc := txd.NewAmount("BTC", 2, 1)
	b.Inc(inc)
	if !b.Get(btc).Equal(txd.NewAmount("BTC", 2, 1)) {
		t.Error()
	}
}

func TestDecBalance(t *testing.T) {
	b := txd.NewBalance()
	dec := txd.NewAmount("BTC", -2, 1)
	b.Inc(dec)
	if !b.Get(btc).Equal(txd.NewAmount("BTC", -2, 1)) {
		t.Error()
	}
}

func TestSendTo(t *testing.T) {
	balances := make(txd.Balances)
	balances[1] = txd.NewBalance()
	amount := txd.NewAmount("BTC", 2, 2)
	balances[1].Inc(amount)

	err := balances.SendTo(1, 2, txd.NewAmount("BTC", 3, 1))
	if err != txd.NoEnoughBalance {
		t.Error()
	}
	if !balances[1].Get(txd.NewCurrency("BTC")).Equal(txd.NewAmount("BTC", 2, 2)) {
		t.Error()
	}
	if !balances[2].Get(txd.NewCurrency("BTC")).Equal(txd.NewAmount("BTC", 0, 0)) {
		t.Error()
	}

	err = balances.SendTo(1, 2, txd.NewAmount("BTC", 1, 2))
	if err != nil {
		t.Error()
	}
	if !balances[1].Get(txd.NewCurrency("BTC")).Equal(txd.NewAmount("BTC", 1, 2)) {
		t.Error()
	}

	if !balances[2].Get(txd.NewCurrency("BTC")).Equal(txd.NewAmount("BTC", 1, 2)) {
		t.Error()
	}
}
