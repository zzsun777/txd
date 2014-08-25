package txd

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
	TX_SEND_TO = iota
	TX_OFFER_CREATE
	TX_OFFER_CANCEL
)
