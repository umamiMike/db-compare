package badger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStorage(t *testing.T) {
	storageDir := "/tmp/badger"
	badger := NewBadger(storageDir)
	fmt.Printf("%#v", badger)
	_, err := NewStorage(storageDir)
	require.NoError(t, err)

}
