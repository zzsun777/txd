package plaintext_test

import (
	"bytes"
	"github.com/aisk/txd/plaintext"
	"os"
	"path"
	"strconv"
	"testing"
	"time"
)

var (
	filePath = path.Join(
		os.TempDir(),
		strconv.Itoa(int(time.Now().Unix()))+"_store.txt",
	)
	datas = [][]byte{
		[]byte("foo"),
		[]byte("bar"),
		[]byte("baz"),
	}
)

func check(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestNew(t *testing.T) {
	store, err := plaintext.New(filePath)
	check(t, err)
	check(t, store.Close())
}

func TestAdd(t *testing.T) {
	store, err := plaintext.New(filePath)
	check(t, err)

	for i := 0; i < len(datas); i++ {
		_, err = store.Add([]byte(datas[i]))
		check(t, err)
	}
	check(t, store.Close())
}

func TestGetAll(t *testing.T) {
	store, err := plaintext.New(filePath)
	check(t, err)
	retrieved, err := store.GetAll()
	check(t, err)
	if len(retrieved) != len(datas) {
		t.Fatalf("data length invalid, except %d, got %d", len(datas), len(retrieved))
	}
	for i := 0; i < len(datas); i++ {
		if !bytes.Equal(datas[i], retrieved[i]) {
			t.Fatalf("data not equal, %v => %v", datas[i], retrieved[i])
		}
	}

	check(t, store.Close())
}
