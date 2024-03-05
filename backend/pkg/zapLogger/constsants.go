package zapLogger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

/**
 * @Author: shenfz
 * @Email: 1328919715@qq.com
 * @Date: 2021/12/27 16:15
 * @Description:
 */

// =========================================
const (
	MESSAGE_NAME_KEY   = "msg"   // 信息命名字段名 eg: msg="xxx"
	LEVEL_NAME_KEY     = "level" // 等级命名字段名 eg: level=info
	TIME_NAME_KEY      = "time"  // timestamp 时间命名字段名
	CODE_LOCALTION_KEY = "at"    // 代码位置
)

func EncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func EncodeDuration(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendInt64(int64(d) / 1000000)
}
