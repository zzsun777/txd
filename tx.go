package txd

import (
	"speter.net/go/exp/math/dec/inf"
)

type Currency [3]byte

type Amount struct {
	Value    *inf.Dec
	Currency Currency
}

type Send struct {
	CreatedAt int64
	From      int64
	To        int64
	Amount    Amount
}

type OfferCreate struct {
}

type OfferCancel struct {
}

const (
	TX_SEND = iota
	TX_OFFER_CREATE
	TX_OFFER_CANCEL
)
