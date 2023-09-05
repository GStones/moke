package model

import (
	"github.com/gstones/moke-kit/orm/nosql"
	"github.com/gstones/moke-kit/orm/nosql/diface"
)

type Dao struct {
	nosql.DocumentBase `bson:"-"`
	Data               *Data `bson:"data"`
}

func (d *Dao) Init(id string, doc diface.ICollection) error {
	key, e := NewAuthKey(id)
	if e != nil {
		return e
	}
	d.initData()
	d.DocumentBase.Init(&d.Data, d.clear, doc, key)
	return nil
}

func NewAuthModel(id string, doc diface.ICollection) (*Dao, error) {
	dm := &Dao{}
	if err := dm.Init(id, doc); err != nil {
		return nil, err
	}
	return dm, nil
}
