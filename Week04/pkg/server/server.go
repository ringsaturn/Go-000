package server

import (
	goContext "context"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/ringsaturn/Go-000/Week04/pkg/biz"
)

// Server define
type Server struct {
	B   *biz.Biz
	app *iris.Application
}

// NewServer will init server
func NewServer(b *biz.Biz) (*Server, error) {
	server := &Server{
		B:   b,
		app: iris.New(),
	}
	server.RegisterRoutes()
	return server, nil
}

// ping for health check
func ping(ctx iris.Context) {
	ctx.Writef("pong")
}

func userProblem(err error) iris.Problem {
	return iris.NewProblem().
		Type("/users").
		Title("Product validation problem").
		Detail(fmt.Sprintf("%v", err)).
		Status(iris.StatusBadRequest)
}

// RegisterRoutes will combine path and func
func (server *Server) RegisterRoutes() {
	server.app.Get("/ping", ping)
	server.app.Get("/users", func(ctx iris.Context) {
		users, err := server.B.GetUsers()
		if err != nil {
			ctx.Problem(userProblem(err), iris.ProblemOptions{
				JSON: iris.JSON{
					Indent: "  ",
				},
				RetryAfter: 400,
			})
		}
		ctx.JSON(users)
	})
}

// Listen will run
func (server *Server) Listen(serverHost string) error {
	err := server.app.Listen(serverHost)
	return err
}

// Shutdown will stop server
func (server *Server) Shutdown() error {
	ctx, cancel := goContext.WithTimeout(goContext.Background(), 20*time.Second)
	defer cancel()
	server.app.Shutdown(ctx)
	return nil
}
