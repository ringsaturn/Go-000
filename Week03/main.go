// week03 -> concurrency
// Question
// > 基于 errgroup 实现一个 http server 的启动和关闭 ，
// 以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type appHandler struct {
	content string
}

func (aH *appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}

type debugHandler struct{}

func (dH *debugHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pprof.Index(w, r)
}

// 业务逻辑 API
func newAppServer() *http.Server {
	// h1 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// }
	// h2 := func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #2!\n")
	// }

	// http.HandleFunc("/", h1)
	// http.HandleFunc("/endpoint", h2)
	// err := http.ListenAndServe(":8080", nil)
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: &appHandler{"hello"},
	}
	return server
}

// debug API
func newDebugServer() *http.Server {
	// err := http.ListenAndServe("localhost:6060", nil)
	server := &http.Server{
		Addr:    "localhost:6060",
		Handler: &debugHandler{},
	}
	return server
}

func app(shutdownCh chan bool, closed chan struct{}) {
	g := new(errgroup.Group)
	appServer := newAppServer()
	debugServer := newDebugServer()

	g.Go(func() error {
		if err := appServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		if err := debugServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})
	go func() {
		<-shutdownCh
		appServer.Shutdown(context.Background())
		debugServer.Shutdown(context.Background())
		close(closed)
	}()
	if err := g.Wait(); err != nil {
		appServer.Shutdown(context.Background())
		debugServer.Shutdown(context.Background())
		close(closed)
	}
}

func main() {
	log.Println("current PID", os.Getpid())
	quit := make(chan os.Signal, 1)
	shutdown := make(chan bool)
	closed := make(chan struct{})
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1)

	go app(shutdown, closed)

	select {
	case <-quit:
		log.Println("receive your info, shutting down")
		shutdown <- true
	}

	<-closed
	log.Println("end")
}
