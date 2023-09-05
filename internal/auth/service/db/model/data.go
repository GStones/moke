package model

import "moke/internal/auth/utils"

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

func (d *Dao) initDefaultData() {
	data := createData()
	data.Uid = utils.GenerateUUID()
}

func (d *Dao) clear() {
	d.Data = nil
}

func (d *Dao) GetUid() string {
	return d.Data.Uid
}
