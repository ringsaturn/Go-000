package biz

import (
	"github.com/pkg/errors"
	"github.com/ringsaturn/GO-000/Week02/code"
	"github.com/ringsaturn/GO-000/Week02/dao"
	"github.com/ringsaturn/GO-000/Week02/model"
)

// Movie biz
func Movie(name string) (*model.Movie, error) {
	movie, err := dao.Movie(name)
	if errors.Is(err, code.ErrNotFound) {
		return nil, errors.Wrapf(code.ErrNotFound, "query err %v", err)
	}
	if err != nil {
		return nil, errors.Wrapf(code.ErrUnknown, "original err %v", err)
	}
	return movie, nil
}
