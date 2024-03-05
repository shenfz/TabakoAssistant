package zapLogger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

/**
 * @Author: shenfz
 * @Email: 1328919715@qq.com
 * @Date: 2022/07/12 9:57
 * @Description:
 */

var defaultLogger LoggerXC
var defaultConf LoggerSetting

type LoggerXC struct {
	*zap.Logger
}

type LoggerSetting struct {
	debugMode     bool
	prefix        string        // 日志名  eg: AppGBR
	rotationTime  time.Duration // 日志分割时间 default 24 * time.Hour
	maxLogKeeping time.Duration // 最长保存日志时间 default 7*24*time.Hour
	separateLevel zapcore.Level // 根据日志等级分文件 >=
	persist       bool          // 日志保存
	zapcore.EncoderConfig
}

func init() {
	defaultConf = LoggerSetting{
		debugMode:     false,
		prefix:        "zap",
		rotationTime:  24 * time.Hour,
		maxLogKeeping: 7 * 24 * time.Hour,
		separateLevel: zapcore.ErrorLevel,
		persist:       false,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     MESSAGE_NAME_KEY, // 信息Tag
			LevelKey:       LEVEL_NAME_KEY,   // 等级Tag
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			TimeKey:        TIME_NAME_KEY,              // timestamp 时间Tag
			EncodeTime:     EncodeTime,                 // 时间的打印格式
			CallerKey:      CODE_LOCALTION_KEY,         // 编码位置命名字段名
			EncodeCaller:   zapcore.ShortCallerEncoder, //长短编码 此处是短编码
			EncodeDuration: EncodeDuration,
		},
	}
	if len(os.Args) > 1 {
		defaultConf.prefix = os.Args[0]
	}

	encoder := zapcore.NewConsoleEncoder(defaultConf.EncoderConfig)
	core := zapcore.NewTee(
		// 保存于磁盘文件
		// zapcore.NewCore(encoder, zapcore.AddSync(getWriter("./"+Default_Log_Name+".log",Default_Log_keep,Deafult_Log_RotationTime)), defaultLevel),
		// 控制台打印
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.LevelEnablerFunc(func(lvl zapcore.Level) bool { return true })),
	)
	// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
	defaultLogger.Logger = zap.New(core, zap.AddCaller())
}

func GetDefaultLogConfig() *LoggerSetting {
	return &defaultConf
}

func (s *LoggerSetting) GetPrefix() string {
	return s.prefix
}

func (s *LoggerSetting) IsDebugMode() bool {
	return s.debugMode
}

func (s *LoggerSetting) GetRotationTime() time.Duration {
	return s.rotationTime
}

func (s *LoggerSetting) GetMaxLogKeeping() time.Duration {
	return s.maxLogKeeping
}

func (s *LoggerSetting) GetSeparateLevel() zapcore.Level {
	return s.separateLevel
}

func (s *LoggerSetting) IsPersist() bool {
	return s.persist
}
