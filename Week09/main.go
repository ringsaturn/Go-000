package main

import (
	"bufio"
	"log"
	"net"
)

func writer(ch chan net.Conn) {
	conn := <-ch
	wr := bufio.NewWriter(conn)
	wr.WriteString("hello")
	wr.Flush()
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	ch := make(chan net.Conn)
	ch <- conn
	<-ch
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		go handleConn(conn)
	}

}
