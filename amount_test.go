package txd_test

import (
	"github.com/aisk/txd"
	"testing"
)

func TestNewAmount(t *testing.T) {
	amount := txd.NewAmount("BTC", 123, 2)
	if amount.String() != "1.23BTC" {
		t.Error()
	}
	amount = txd.NewAmount("BTC", -123, 2)
	if amount.String() != "-1.23BTC" {
		t.Error()
	}
}

func TestAmountEqual(t *testing.T) {
	this := txd.NewAmount("BTC", 1, 1)
	that := txd.NewAmount("BTC", 1, 1)

	if !this.Equal(that) {
		t.Error()
	}

	if !this.Equal(txd.NewAmount("BTC", 10, 2)) {
		t.Error()
	}

	if this.Equal(txd.NewAmount("DOG", 1, 1)) {
		t.Error()
	}
}

func TestAmountNeg(t *testing.T) {
	a := txd.NewAmount("BTC", 12, 3)
	a.Neg()
	if !a.Equal(txd.NewAmount("BTC", -12, 3)) {
		t.Error()
	}
	a.Neg()
	if !a.Equal(txd.NewAmount("BTC", 12, 3)) {
		t.Error()
	}
}

func TestAmountCmp(t *testing.T) {
	a := txd.NewAmount("BTC", 23, 4)
	b := txd.NewAmount("BTC", 22, 4)
	if !(a.Cmp(b) > 0) {
		t.Error()
	}
	if !(a.Cmp(a) == 0) {
		t.Error()
	}
}
