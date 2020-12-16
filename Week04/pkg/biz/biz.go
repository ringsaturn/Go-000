package biz

import (
	"github.com/ringsaturn/Go-000/Week04/pkg/dao"
	"github.com/ringsaturn/Go-000/Week04/pkg/model"
)

// Biz define
type Biz struct {
	D *dao.Dao
}

// NewBiz will init
func NewBiz(d *dao.Dao) (*Biz, error) {
	return &Biz{
		D: d,
	}, nil
}

// GetUsers will return users
func (biz *Biz) GetUsers() ([]model.User, error) {
	return biz.D.QueryUser()
}
