// Package week02 -> error
// Question
// 	我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，
// 	是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
package main

import (
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"github.com/ringsaturn/Go-000/Week02/biz"
	"github.com/ringsaturn/Go-000/Week02/code"
)

func movieService(ctx iris.Context) {
	movieInfo, err := biz.Movie(ctx.URLParam("name"))
	if err != nil {
		statusCode := 500
		if errors.Is(err, code.ErrNotFound) {
			statusCode = 404
		}
		if errors.Is(err, code.ErrUnknown) {
			statusCode = 500
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
