package db

import (
	"fmt"
	"testing"
)

func TestNewBadger(t *testing.T) {
	storageDir := "/tmp/badger"
	badger := NewBadger(storageDir)
	fmt.Printf("%#v", badger)

}
