package db

import (
	"errors"
	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"go.uber.org/zap"
	"moke/internal/auth/service/db/model"
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

func (d *Database) NewBuddyQueue(id string) (*model.Dao, error) {
	bq := new(model.Dao)
	err := bq.Init(id, d.coll)
	if err != nil {
		return nil, err
	}
	return bq, nil
}

func (d *Database) CreateBuddyQueue(id string) error {
	if bq, err := d.NewBuddyQueue(id); err != nil {
		return err
	} else if err = bq.Create(); err != nil {
		return err
	}
	return nil
}

func (d *Database) LoadOrCreateBuddyQueue(id string) (bq *model.Dao, err error) {
	if bq, err = d.NewBuddyQueue(id); err != nil {
		return
	} else if err = bq.Load(); errors.Is(err, nerrors.ErrKeyNotFound) {
		if bq, err = d.NewBuddyQueue(id); err != nil {
			return
		} else if err := bq.InitDefault(); err != nil {
			return nil, err
		} else if err = bq.Create(); err != nil {
			err = bq.Load()
		}
	}
	return
}
