package memdb

import (
	"testing"
	"time"
)

func TestMemDB(t *testing.T) {
	m := New()
	m.Write("abcd")

	if len(m.DBObjs) != 1 {
		t.Errorf("Expected length to be 1, got %v.", len(m.DBObjs))
	}

}

func TestQuery(t *testing.T) {
	m := New()
	m.Write("abcd")

	currTime := time.Now().UTC()
	sTime := currTime.Unix()
	eTime := currTime.Add(time.Second).Unix()
	res := m.Query(sTime, eTime)
	if len(res) != 1 {
		t.Errorf("Expected query res to be 1, got %v.", len(res))
	}
}
