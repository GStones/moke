package db

import (
	"errors"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"go.uber.org/zap"
)

type Database struct {
	logger *zap.Logger
	coll   diface.ICollection
}

func OpenDatabase(l *zap.Logger, coll diface.ICollection) *Database {
	return &Database{
		logger: l,
		coll:   coll,
	}
}

func (d *Database) NewBuddyQueue(appId string, id string) (*BuddyQueue, error) {
	bq := new(BuddyQueue)
	err := bq.init(appId, id, d.coll)
	if err != nil {
		return nil, err
	}
	return bq, nil
}

func (d *Database) ContainsBuddyQueue(appId string, id string) (bool, error) {
	if key, e := newBuddyQueueKey(appId, id); e != nil {
		return false, e
	} else {
		return d.coll.Contains(key)
	}
}

func (d *Database) IncrBuddyReward(appId string, id, path string, value int32) error {
	if key, e := newBuddyQueueKey(appId, id); e != nil {
		return e
	} else {
		return d.coll.Incr(key, path, value)
	}
}

func (d *Database) PushBack(appId string, id, path string, profileId string) error {
	if key, e := newBuddyQueueKey(appId, id); e != nil {
		return e
	} else {
		return d.coll.PushBack(key, path, profileId)
	}
}

func (d *Database) CreateBuddyQueue(appId, id string) error {
	if bq, err := d.NewBuddyQueue(appId, id); err != nil {
		return err
	} else if err = bq.Create(); err != nil {
		return err
	}
	return nil
}

func (d *Database) LoadOrCreateBuddyQueue(appId string, id string) (bq *BuddyQueue, err error) {
	if bq, err = d.NewBuddyQueue(appId, id); err != nil {
		return
	} else if err = bq.Load(); errors.Is(err, nerrors.ErrKeyNotFound) {
		if bq, err = d.NewBuddyQueue(appId, id); err != nil {
			return
		} else if err := bq.InitDefault(); err != nil {
			return nil, err
		} else if err = bq.Create(); err != nil {
			err = bq.Load()
		}
	}
	// Even if we know the BuddyQueue is on the latest version, we should still run it through fixups
	if err == nil {
		err = d.FixupBuddyQueue(bq)
	}
	return
}
func (d *Database) LoadBuddyQueue(appId string, id string) (bq *BuddyQueue, err error) {
	if bq, err = d.NewBuddyQueue(appId, id); err != nil {
		return
	} else if err = bq.Load(); err != nil {
		return
	}
	// Even if we know the BuddyQueue is on the latest version, we should still run it through fixups
	if err == nil {
		err = d.FixupBuddyQueue(bq)
	}
	return
}
