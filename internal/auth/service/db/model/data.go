package model

import (
	"moke/internal/auth/service/utils"
)

type Data struct {
	Uid string `json:"uid" bson:"uid"`
}

func createData() *Data {
	return &Data{}
}

func (d *Dao) initData() {
	data := createData()
	d.Data = data
}

func (d *Dao) InitDefault() error {
	data := createData()
	data.Uid = utils.GenerateUUID()
	d.Data = data
	return nil
}

func (d *Dao) clear() {
	d.Data = nil
}

func (d *Dao) GetUid() string {
	return d.Data.Uid
}
