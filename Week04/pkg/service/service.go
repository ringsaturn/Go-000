package service

import (
	"github.com/ringsaturn/Go-000/Week04/pkg/server"
)

// Service define
type Service struct {
	Server *server.Server
}

// NewService will init new srv
func NewService(server *server.Server) (*Service, error) {
	srv := &Service{
		Server: server,
	}
	return srv, nil
}

// Start define
func (srv *Service) Start(serverHost string) {
	err := srv.Server.Listen(serverHost)
	if err != nil {
		panic(err)
	}
}

// Register 注册
func (srv *Service) Register() error {
	return nil
}

// Unregister 取消注册
func (srv *Service) Unregister() error {
	return nil
}

// Shutdown define
func (srv *Service) Shutdown() {
	// 关闭服务器
	srv.Server.Shutdown()
	// 关闭数据库
	_ = srv.Server.B.D.UserDB.Close()
	return
}
