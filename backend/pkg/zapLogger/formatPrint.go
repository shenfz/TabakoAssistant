package zapLogger

import (
	"fmt"
)

/**
 * @Author: shenfz
 * @Author: 1328919715@qq.com
 * @Date: 2021/3/31 17:21
 * @Desc:
 */

func (z *LoggerXC) Debugf(format string, strs ...interface{}) {
	z.Debug(fmt.Sprintf(format, strs...))
}

func (z *LoggerXC) Infof(format string, strs ...interface{}) {
	z.Info(fmt.Sprintf(format, strs...))
}

func (z *LoggerXC) Warnf(format string, strs ...interface{}) {
	z.Warn(fmt.Sprintf(format, strs...))
}

func (z *LoggerXC) Errorf(format string, strs ...interface{}) {
	z.Error(fmt.Sprintf(format, strs...))
}

func (z *LoggerXC) Fatalf(format string, strs ...interface{}) {
	z.Fatal(fmt.Sprintf(format, strs...))
}
