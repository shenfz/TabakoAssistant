package zapLogger

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GetGlobalLogger() *LoggerXC {
	return &defaultLogger
}

func NewLogger(opts ...LogOptFunc) *LoggerXC {
	var (
		conf = GetDefaultLogConfig()
	)

	if len(opts) == 0 {
		return GetGlobalLogger()
	}

	for i := range opts {
		opts[i](conf)
	}

	encoder := zapcore.NewConsoleEncoder(conf.EncoderConfig)
	// 实现两个判断日志等级的interface 分割日志
	splitBefore := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		// return lvl < conf.GetSeparateLevel()
		return true
	})
	splitAfter := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl > conf.GetSeparateLevel()
	})

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(getSplitBeforeWriter()), splitBefore),
		zapcore.NewCore(encoder, zapcore.AddSync(getSplitAfterWriter()), splitAfter),
		getConsoleStdoutWriter(encoder))

	defaultLogger.Logger = zap.New(core, zap.AddCaller()) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	return &defaultLogger
}

func getSplitBeforeWriter() io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	conf := GetDefaultLogConfig()
	hook, err := rotatelogs.New(
		//fmt.Sprintf("%s%s%s", GetRootDirPath(), string(filepath.Separator), conf.GetPrefix())+".all.%Y%m%d%H",
		GetRootDirPath()+conf.GetPrefix()+".all.%Y%m%d%H",
		//	rotatelogs.WithLinkName("./logs"+conf.GetPrefix()+".all"),
		rotatelogs.WithMaxAge(conf.GetMaxLogKeeping()),
		rotatelogs.WithRotationTime(conf.GetRotationTime()),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func GetRootDirPath() string {
	dir, _ := os.Getwd()
	rootDir := dir[:strings.LastIndex(dir, string(filepath.Separator))]
	return rootDir + string(filepath.Separator)
}

func getSplitAfterWriter() io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	conf := GetDefaultLogConfig()
	hook, err := rotatelogs.New(
		//fmt.Sprintf("%s%s%s", GetRootDirPath(), string(filepath.Separator), conf.GetPrefix())+".all.%Y%m%d%H",
		// rotatelogs.WithLinkName("./logs/"+conf.GetPrefix()+".err"),
		GetRootDirPath()+conf.GetPrefix()+".all.%Y%m%d%H",
		rotatelogs.WithMaxAge(conf.GetMaxLogKeeping()),
		rotatelogs.WithRotationTime(conf.GetRotationTime()),
	)
	if err != nil {
		panic(err)
	}
	return hook
}

func getConsoleStdoutWriter(encoder zapcore.Encoder) zapcore.Core {
	return zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	}))
}
