package logs

import (
	"fmt"
	"github.com/asd864613087/logs/consts"
	"github.com/asd864613087/logs/utils"
	"net"
	"time"
)

var (
	consoleProvider = &ConsoleProvide{}
	logAgentProvider = &LogAgentProvider{}
	fileProvider = &FileProvider{}
)

type Provider interface {
	WriteMsg(msg []byte)
	Init()
	SetLevel(level int)
	Flush()
	Finish()
}

type ConsoleProvide struct {
	level int
}

func (c ConsoleProvide) WriteMsg(msg []byte) {
	// TODO：考虑重定向问题
	fmt.Println(string(msg))
}

func (c ConsoleProvide) Init() {
	c.level = consts.DEFAULT_CONSOLE_PROVIDER_LEVEL
	return
}

func (c ConsoleProvide) SetLevel(level int) {
	c.level = level
}

func (c ConsoleProvide) Flush() {
	return
}


func (c ConsoleProvide) Finish() {
	return
}

type LogAgentProvider struct {
	level int
	maxBufSize int
	curBufPos int

	buf []byte
	conn net.Conn
	signal chan int
}

func (l *LogAgentProvider) WriteMsg(msg []byte) {
	// TODO: 如何处理界限问题，logAgent需要对消息进行marshol
	//if l.curBufPos + len(msg) > l.maxBufSize {
	//	l.Flush()
	//}
	//
	l.buf= append(l.buf, msg...)
	l.Flush()
}

func (l *LogAgentProvider) Init() {
	file := utils.GetUnixPath()
	// os.Remove(file)

	// TODO: 修改这部分初始化逻辑的位置
	addr, err := net.ResolveUnixAddr("unix", file)
	if err != nil {
		fmt.Printf("[LogAgentProvider.Init] ResolveUnixAddr Failed: err = %s", err)
		return
	}

	for i := 0; i < 30; i++ {
		conn, err := net.DialUnix("unix", nil, addr)
		if err != nil {
			fmt.Printf("[LogAgentProvider.Init] DialUnix Failed: err = %s", err)
			time.Sleep( 1 * time.Second)
			continue
		}
		l.conn = conn
		break
	}

	if l.conn == nil {
		panic("[LogAgentProvider.Init] DialUnix Failed")
	}

	l.level = consts.DEFAULT_LOGAGENT_PROVIDER_LEVEL
	l.buf = make([]byte, 128)
	l.signal = make(chan int)
}

func (l *LogAgentProvider) SetLevel(level int) {
	l.level = level
}

func (l *LogAgentProvider) Flush() {
	// 再打包， 加入自定义协议头
	buf := utils.PackStreamData(l.buf)

	// 将缓存内数据推到socket
	_, err := l.conn.Write(buf)
	if err != nil {
		fmt.Printf("[LogAgentProvider.Flush] Flush: err = %s", err)
		return
	}

	// 清空缓存
	l.buf = l.buf[:0]
}


func (l *LogAgentProvider) Finish() {
	err := l.conn.Close()
	if err != nil {
		fmt.Printf("[LogAgentProvider.Finish] conn.Close Failed: err = %s", err)
		return
	}
}

type FileProvider struct {

}

func (f FileProvider) WriteMsg(msg []byte) {
	panic("implement me")
}

func (f FileProvider) Init() {
	panic("implement me")
}

func (f FileProvider) SetLevel(level int) {
	panic("implement me")
}

func (f FileProvider) Flush() {
	panic("implement me")
}

func (f FileProvider) Finish() {
	panic("implement me")
}



