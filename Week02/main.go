// Package week02 -> error
// Question
// 	我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 	是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
package main

import (
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
)

// ErrNotFound error
var ErrNotFound = errors.New("NotFound")

// Movie define info for a movie
type Movie struct {
	Name   string `json:"name"`
	EnName string `json:"enname"`
}

func dao(name string) (*Movie, error) {
	if name != "蜜桃成熟时33D" {
		return nil, errors.WithMessage(ErrNotFound, name)
	}
	movie := &Movie{
		Name:   name,
		EnName: "The 33D Invader",
	}
	return movie, nil
}

func biz(name string) (*Movie, error) {
	movie, err := dao(name)
	if err != nil {
		return nil, err
	}
	return movie, nil
}

func movieService(ctx iris.Context) {
	movieInfo, err := biz(ctx.URLParam("name"))
	if err != nil {
		statusCode := 400
		if errors.Is(err, errors.New("NotFound")) {
			statusCode = 404
		}
		ctx.StopWithProblem(statusCode, iris.NewProblem().
			Title("FailFetchMovie").DetailErr(err))
		return
	}
	ctx.JSON(iris.Map{
		"status": "ok",
		"data":   movieInfo,
	})
	return
}

func main() {
	app := iris.New()
	app.Get("/movie", movieService)
	app.Listen("localhost:8080")
}
