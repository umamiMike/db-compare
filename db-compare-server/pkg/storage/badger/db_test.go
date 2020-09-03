package db

import (
	"fmt"
	"testing"
)

func TestNewBadger(t *testing.T) {
	storageDir := "/tmp/badger"
	badger := NewBadger(storageDir)
	fmt.Printf("%#v", badger)
	//assert.Equal(t, storageDir, badger.opts.Dir, "new badger should return Badger with values, opts containaing storageDir")

}
