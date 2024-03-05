package zapLogger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

/**
 * @Author: shenfz
 * @Email: 1328919715@qq.com
 * @Date: 2022/07/12 9:57
 * @Description:
 */

type LogOptFunc func(c *LoggerSetting)

func SetDebugMode() LogOptFunc {
	return func(c *LoggerSetting) {
		c.debugMode = true
		c.EncodeCaller = zapcore.FullCallerEncoder // 长编码
	}
}

func SetPrefix(prefix string) LogOptFunc {
	return func(c *LoggerSetting) {
		if prefix != "" {
			c.prefix = prefix
		}
	}
}

func SetPersist() LogOptFunc {
	return func(c *LoggerSetting) {
		c.persist = true
	}
}

func SetRotationTime(rt time.Duration) LogOptFunc {
	return func(c *LoggerSetting) {
		if rt >= time.Hour {
			c.rotationTime = rt
		}
	}
}

func SetMaxLogKeeping(mk time.Duration) LogOptFunc {
	return func(c *LoggerSetting) {
		c.maxLogKeeping = mk
	}
}

func SetSeparateLevel(sl zapcore.Level) LogOptFunc {
	return func(c *LoggerSetting) {
		c.separateLevel = sl
	}
}
