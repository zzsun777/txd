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
	println(balances)
}
