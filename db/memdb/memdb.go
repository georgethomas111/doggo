package memdb

import "time"

type dbObj struct {
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type MemDB struct {
	DBObjs []dbObj
}

func New() *MemDB {
	return &MemDB{}
}

func (m *MemDB) Write(obj interface{}) {
	o := dbObj{
		Timestamp: time.Now().UTC().Unix(),
		Data:      obj,
	}

	m.DBObjs = append(m.DBObjs, o)
}
