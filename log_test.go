package logs

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestCtxInfo(t *testing.T) {
	go server()

	time.Sleep(2 * time.Second)

	c, err := net.Dial("unix", "C:\\Users\\86461\\echo.sock")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	defaultLogger.Start()
	fmt.Printf("[LogAgentProvider] Initial: provider=%+v", defaultLogger.providerList[1])

	i := 0
	for {
		ctx := context.Background()
		CtxInfo(ctx, "Test Log: %d", i)
		time.Sleep(1 * time.Second)
		i++
		fmt.Println(i)
	}


}

func server()  {
	err := os.Remove("C:\\Users\\86461\\echo.sock")
	if err != nil {
		fmt.Println(err)
	}

	l, err := net.Listen("unix", "C:\\Users\\86461\\echo.sock")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	for {
		fd, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(fd)
	}
}

func echoServer(c net.Conn) {
	for {
		buf := make([]byte, 512)
		nr, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[0:nr]
		println("Server got:", string(data))
		_, err = c.Write(data)
		if err != nil {
			log.Fatal("Write: ", err)
		}
	}
}
