package plaintext

import (
	"bufio"
	"os"
)

type Store struct {
	f    *os.File
	rdwr *bufio.ReadWriter
}

func New(path string) (*Store, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	rd := bufio.NewReader(f)
	wr := bufio.NewWriter(f)
	return &Store{f, bufio.NewReadWriter(rd, wr)}, nil
}

func (store *Store) Add(data []byte) (int64, error) {
	if _, err := store.f.Seek(0, os.SEEK_END); err != nil {
		return 0, err
	}

	if _, err := store.rdwr.Write(data); err != nil {
		return 0, err
	}
	if _, err := store.rdwr.Write([]byte("\r\n")); err != nil {
		return 0, err
	}
	if err := store.rdwr.Flush(); err != nil {
		return 0, err
	}
	return 1, nil
}

func (store *Store) Get(idx int64) ([]byte, error) {
	panic("not implemented")
}

func (store *Store) GetAll() ([][]byte, error) {
	if _, err := store.f.Seek(0, os.SEEK_SET); err != nil {
		return nil, err
	}

	var datas [][]byte
	scanner := bufio.NewScanner(store.rdwr)
	for scanner.Scan() {
		datas = append(datas, scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return datas, nil
}

func (store *Store) Close() error {
	return store.f.Close()
}
