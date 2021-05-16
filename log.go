package logs

import (
	"context"
)

func CtxDebug(ctx context.Context, fmtStr string, value ...interface{})  {
	defaultLogger.ctxInfo(ctx, fmtStr, value)
}

func CtxInfo(ctx context.Context, fmtStr string, value ...interface{}) {
	defaultLogger.ctxInfo(ctx, fmtStr, value)
}

func CtxWarn(ctx context.Context, fmtStr string, value ...interface{}) {
	defaultLogger.ctxInfo(ctx, fmtStr, value)
}

func CtxError(ctx context.Context, fmtStr string, value ...interface{}) {
	defaultLogger.ctxInfo(ctx, fmtStr, value)
}

func Debug(fmtStr string, value ...interface{}){
	defaultLogger.Debug(fmtStr, value)
}

func Info(fmtStr string, value ...interface{}){
	defaultLogger.Info(fmtStr, value)
}

func Warn(fmtStr string, value ...interface{}){
	defaultLogger.Warn(fmtStr, value)
}

func Error(fmtStr string, value ...interface{}){
	defaultLogger.Error(fmtStr, value)
}





