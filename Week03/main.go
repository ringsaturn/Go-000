// week03 -> concurrency
// Question
// > 基于 errgroup 实现一个 http server 的启动和关闭 ，
// 以及 linux signal 信号的注册和处理，要保证能够 一个退出，全部注销退出。
package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
)

// 业务逻辑 API
func appServer() error {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	err := http.ListenAndServe(":8080", nil)
	return err
}

// debug API
func debugServer() error {
	err := http.ListenAndServe("localhost:6060", nil)
	return err
}

func main() {
	log.Println("current PID", os.Getpid())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1)
	<-quit
}
