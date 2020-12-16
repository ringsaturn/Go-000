// 按照自己的构想，写一个项目满足基本的目录结构和工程，
// 代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，信号处理，使用 Wire 构建依赖。
// 可以使用自己熟悉的框架。
package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ringsaturn/Go-000/Week04/pkg/biz"
	"github.com/ringsaturn/Go-000/Week04/pkg/dao"
	"github.com/ringsaturn/Go-000/Week04/pkg/fakedb"
	"github.com/ringsaturn/Go-000/Week04/pkg/server"
	"github.com/ringsaturn/Go-000/Week04/pkg/service"
)

// ServerHost for server
const ServerHost = "localhost:8888"

// DBHost for db
const DBHost = "localhost:666"

func main() {
	fakeDB, err := fakedb.NewFakeDB(DBHost)
	if err != nil {
		panic(err)
	}
	d, err := dao.NewDao(fakeDB)
	if err != nil {
		panic(err)
	}
	myBiz, err := biz.NewBiz(d)
	if err != nil {
		panic(err)
	}
	myServer, err := server.NewServer(myBiz)
	if err != nil {
		panic(err)
	}
	srv, err := service.NewService(myServer)
	if err != nil {
		panic(err)
	}
	go srv.Start(ServerHost)
	// 等待启动完成
	for i := 0; i <= 10; i++ {
		resp, err := http.Get(ServerHost + "/ping")

		if err != nil {
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			srv.Register()
			break
		}
		time.Sleep(1 * time.Second)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1)

	select {
	case <-quit:
		srv.Unregister()
		srv.Shutdown()
	}
}
