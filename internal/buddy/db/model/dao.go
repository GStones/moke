package model

import (
	"github.com/gstones/moke-kit/orm/nerrors"
	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
	"moke/internal/buddy/db/model/data"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *data.BuddyQueue `bson:"data"`
}

func (b *Dao) Init(id string, ros diface.ICollection) error {
	if ros == nil {
		return nerrors.ErrDocumentStoreIsNil
	}
	k, e := NewBuddyQueueKey(id)
	if e != nil {
		return e
	}
	b.Data = data.NewBuddyQueue(id)
	b.DocumentBase.Init(&b.Data, b.clear, ros, k)
	return nil
}

func (b *Dao) clear() {
	b.Data.Clear()
}

func (b *Dao) InitDefault() error {
	return nil
}
