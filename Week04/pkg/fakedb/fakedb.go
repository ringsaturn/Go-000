package fakedb

import (
	"errors"
	"sync"
)

// FakeDB define
type FakeDB struct {
	Host  string
	Ready bool
	conn  int16
	mu    sync.Mutex
}

// ErrAlreadyConn will raise for dup Conn calls
var ErrAlreadyConn = errors.New("AlreadyConn")

// ErrAlreadyClose will raise for dup Close calls
var ErrAlreadyClose = errors.New("AlreadyClose")

// NewFakeDB will do init
func NewFakeDB(Host string) (*FakeDB, error) {
	db := &FakeDB{
		Host: Host,
	}
	err := db.Conn()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Conn define
func (db *FakeDB) Conn() error {
	if db.conn == 1 {
		return ErrAlreadyConn
	}
	db.mu.Lock()
	db.conn++
	db.Ready = true
	db.mu.Unlock()
	return nil
}

// GetMulti define
func (db *FakeDB) GetMulti() ([]interface{}, error) {
	return nil, nil
}

// Close define
func (db *FakeDB) Close() error {
	if db.conn == 0 {
		return ErrAlreadyConn
	}
	db.mu.Lock()
	db.conn--
	db.Ready = false
	db.mu.Unlock()
	return nil
}
