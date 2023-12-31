package db

import (
	"errors"

	"go.uber.org/zap"

	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql/diface"

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

func (db *Database) LoadOrCreateUid(id string) (*model.Dao, error) {
	if dm, err := model.NewAuthModel(id, db.coll); err != nil {
		return nil, err
	} else if err = dm.Load(); errors.Is(err, nerrors.ErrNotFound) {
		if dm, err = model.NewAuthModel(id, db.coll); err != nil {
			return nil, err
		} else if err := dm.InitDefault(); err != nil {
			return nil, err
		} else if err = dm.Create(); err != nil {
			if err = dm.Load(); err != nil {
				return nil, err
			} else {
				return dm, nil
			}
		} else {
			return dm, nil
		}
	} else if err != nil {
		return nil, err
	} else {
		return dm, nil
	}
}
