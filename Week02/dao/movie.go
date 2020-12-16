package dao

import (
	"errors"

	"github.com/ringsaturn/Go-000/Week02/model"
)

// ErrNotFound error
var ErrNotFound = errors.New("NotFound")

func Movie(name string) (*model.Movie, error) {
	if name != "蜜桃成熟时33D" {
		return nil, ErrNotFound
	}
	movie := &model.Movie{
		Name:   name,
		EnName: "The 33D Invader",
	}
	return movie, nil
}
