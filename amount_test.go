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
