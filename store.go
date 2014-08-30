package txd

import ()

type Store interface {
	Add([]byte) (int64, error)
	Get(int64) ([]byte, error)
	GetAll() ([][]byte, error)
	Close() error
}
