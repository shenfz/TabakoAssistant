package zapLogger

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/3/11 10:48
 * @Desc:
 */

// 使用默认日志器 只显示终端，不写日志文件
func Test_GetGlobalLogger(t *testing.T) {
	logger := GetGlobalLogger()
	logger.Debug("debug", zap.String("string", "d"))
	logger.Info("info", zap.Int("int", 19))
	logger.Warn("warn", zap.Float64("float", 3.124))
	logger.Error("error", zap.ByteString("byte", []byte("xxxx")))

}

func Test_Fields(t *testing.T) {
	logger := GetGlobalLogger().With(zap.String("Index", "test"), zap.String("Type", "t1"))
	logger.Debug("debug", zap.String("string", "d"))
	logger.Info("info", zap.Int("int", 19))
	logger.Warn("warn", zap.Float64("float", 3.124))
	logger.Error("error", zap.ByteString("byte", []byte("xxxx")))
}

func Test_InitSetting(t *testing.T) {
	logger := NewLogger(SetPersist(), SetDebugMode(), SetPrefix("Hello"))
	logger.Debug("debug", zap.String("string", "d"))
	logger.Info("info", zap.Int("int", 19))
	logger.Warn("warn", zap.Float64("float", 3.124))
	logger.Error("error", zap.ByteString("byte", []byte("xxxx")))
	time.Sleep(2 * time.Second)

	logger = NewLogger(SetPersist(), SetDebugMode(), SetPrefix("3Q"))
	logger.Debug("debug", zap.String("string", "d"))
	logger.Info("info", zap.Int("int", 19))
	logger.Warn("warn", zap.Float64("float", 3.124))
	logger.Error("error", zap.ByteString("byte", []byte("xxxx")))

}
