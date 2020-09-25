package db

import "time"

func (storage *Badger) runStorageGC() {
	timer := time.NewTicker(10 * time.Minute)
	for {
		select {
		case <-timer.C:
			storage.storageGC()
		}
	}
}
func (storage *Badger) storageGC() {
again:
	err := storage.db.RunValueLogGC(0.5)
	if err == nil {
		goto again
	}

}

func (storage *Badger) Close() {
	storage.db.Close()
}
