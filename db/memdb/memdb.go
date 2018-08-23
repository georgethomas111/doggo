package memdb

import (
	"sync"
	"time"
)

type dbObj struct {
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type MemDB struct {
	sync.Mutex
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

func (m *MemDB) Query(sTime int64, eTime int64) []interface{} {
	var res []interface{}
	for _, obj := range m.DBObjs {
		if sTime <= obj.Timestamp && obj.Timestamp <= eTime {
			res = append(res, obj)
		}
	}
	return res
}
