package dao

import (
	"errors"

	"github.com/ringsaturn/Go-000/Week04/pkg/fakedb"
	"github.com/ringsaturn/Go-000/Week04/pkg/model"
)

// Dao define
type Dao struct {
	UserDB *fakedb.FakeDB
}

// ErrInvalidDB will return when db is not ready
var ErrInvalidDB = errors.New("InvalidDB")

// NewDao init
func NewDao(UserDB *fakedb.FakeDB) (*Dao, error) {
	if UserDB.Ready == false {
		return nil, ErrInvalidDB
	}
	return &Dao{
		UserDB: UserDB,
	}, nil
}

// QueryUser define
func (dao *Dao) QueryUser() ([]model.User, error) {
	alice := model.User{
		ID:   1,
		Name: "Alice",
	}
	bob := model.User{
		ID:   2,
		Name: "Bob",
	}
	return []model.User{
		alice, bob,
	}, nil
}
