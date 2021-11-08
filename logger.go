package logs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/asd864613087/logs/consts"
	"github.com/asd864613087/logs/utils"
	"sync"
)

var (
	defaultLogger *Logger
	loggerKeys = []string{"logId", "upstreamPsm","psm"}
)

type Logger struct {
	// 日志等级
	level int

	// 服务标识
	psm string

	// 日志输出途径
	providerList []Provider

	// Chan
	buf    chan []byte
	signal chan int

	// kvs, 需要处理的kv对
	// TODO: 上游psm等
	kvs map[string]interface{}

	// 线程与信号
	threadCnt  int
	WaitGroups sync.WaitGroup //TODO：怎样调用wait
	lock       sync.RWMutex
}

func init() {

	psm := utils.GetPsm()

	// TODO: 修改默认Logger
	defaultLogger = &Logger{
		level: consts.DEFAULT_LOGGER_LOGLEVEL,
		psm:   psm,
		providerList: []Provider{
			consoleProvider,
			logAgentProvider,
		},
		buf:        make(chan []byte, 128),
		signal:     make(chan int, 1),
		kvs:        map[string]interface{}{},
		threadCnt:  3,
		WaitGroups: sync.WaitGroup{},
	}
	for _, item := range loggerKeys {
		defaultLogger.kvs[item] = ""
	}

	for _, p := range defaultLogger.providerList {
		p.Init()
	}

	defaultLogger.Start()
}

func NewLogger(logger *Logger) {
	defaultLogger = logger
}

func (logger *Logger) Start() {
	defer func() {
		// logger.WaitGroups.Wait()
	}()

	for i := 0; i < logger.threadCnt; i++ {
		go logger.run()
		logger.WaitGroups.Add(1)
	}
}

func (logger *Logger) run() {
	defer func() {
		logger.WaitGroups.Done()
	}()

	for {
		select {
		// 信息读入
		case msg := <-logger.buf:
			// 遍历输出途径
			for _, p := range logger.providerList {
				p.WriteMsg(msg)
			}

		// 信号,TODO: 单信号只能停止一个线程
		case sig := <-logger.signal:
			if sig == consts.STOP {
				logger.Finish()
				break
			}
		}
	}
}

func (logger *Logger) Finish() {
	logger.lock.Lock()
	defer func() {
		logger.lock.Unlock()
	}()

	close(logger.signal)

	for _, p := range logger.providerList {
		// TODO: 是否需要最后一次Flush
		p.Flush()

		p.Finish()
	}
}

// 处理ctx中的kv
func (logger *Logger) fmtKvLog(ctx context.Context, fmtStr string, value []interface{}) {
	// 更新kv
	for k := range logger.kvs {
		v := ctx.Value(k)

		logger.kvs[k] = v
	}

	logger.fmtLog(fmtStr, value)
}

// 处理Log
func (logger *Logger) fmtLog(fmtStr string, value []interface{}) {
	fmtStr = fmt.Sprintf(fmtStr, value...)
	w := bytes.NewBuffer([]byte{})

	// 处理前缀
	w.Write([]byte(fmt.Sprintf("psm=%s", logger.psm)))
	w.Write([]byte{' '})
	w.Write([]byte(fmt.Sprintf("date=%s", utils.GetCurrentTime())))
	w.Write([]byte{' '})
	w.Write([]byte(fmt.Sprintf("podName=%s", utils.GetPodName())))
	w.Write([]byte{' '})

	// 处理Kv
	for k, v := range logger.kvs {
		b, err := utils.CtxValueToBytes(v)
		if err != nil {
			fmt.Printf("[fmtLog] Error: %+v \n", err)
		}
		w.Write([]byte(k))
		w.Write([]byte{'='})
		w.Write(b)
		w.Write([]byte{' '})
	}

	// 最后写入raw msg
	w.Write([]byte(fmtStr))

	select {
	case logger.buf <- w.Bytes():
	}

}

func (logger *Logger) ctxDebug(ctx context.Context, fmtStr string, value []interface{}) {
	if logger.level < consts.LOGLEVEL_DEBUG {
		return
	}

	logger.fmtKvLog(ctx, fmtStr, value)
}

func (logger *Logger) ctxInfo(ctx context.Context, fmtStr string, value []interface{}) {
	if logger.level < consts.LOGLEVEL_INFO {
		return
	}

	logger.fmtKvLog(ctx, fmtStr, value)
}

func (logger *Logger) ctxWarn(ctx context.Context, fmtStr string, value []interface{}) {
	if logger.level < consts.LOGLEVEL_WARN {
		return
	}

	logger.fmtKvLog(ctx, fmtStr, value)
}

func (logger *Logger) ctxError(ctx context.Context, fmtStr string, value []interface{}) {
	if logger.level < consts.LOGLEVEL_ERROR {
		return
	}

	logger.fmtKvLog(ctx, fmtStr, value)
}

func (logger *Logger) Debug(fmtStr string, value []interface{})  {
	if logger.level < consts.LOGLEVEL_DEBUG {
		return
	}

	logger.fmtLog(fmtStr, value)
}

func (logger *Logger) Info(fmtStr string, value []interface{})  {
	if logger.level < consts.LOGLEVEL_INFO {
		return
	}

	logger.fmtLog(fmtStr, value)
}

func (logger *Logger) Warn(fmtStr string, value []interface{})  {
	if logger.level < consts.LOGLEVEL_WARN {
		return
	}

	logger.fmtLog(fmtStr, value)
}

func (logger *Logger) Error(fmtStr string, value []interface{})  {
	if logger.level < consts.LOGLEVEL_ERROR {
		return
	}

	logger.fmtLog(fmtStr, value)
}



