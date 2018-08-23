package memdb

import "testing"

func TestMemDB(t *testing.T) {
	m := New()
	m.Write("abcd")

	if len(m.DBObjs) != 1 {
		t.Errorf("Expected length to be 1, got %v.", len(m.DBObjs))
	}

}
