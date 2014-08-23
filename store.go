package txd

import (
	"github.com/aisk/txd/plaintext"
)

type Store interface {
	Add([]byte) (int64, error)
	Get(int64) ([]byte, error)
	GetAll() ([][]byte, error)
	Close() error
}

var store Store

func init() {
	var err error
	store, err = plaintext.New("/tmp/test.txt")
	if err != nil {
		panic(err)
	}
}
