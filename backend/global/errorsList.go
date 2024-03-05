package global

import "errors"

/**
 * @Author shenfz
 * @Date 2024/3/5 15:57
 * @Email 1328919715@qq.com
 * @Description:
 **/

var (
	AutoCloseWssConn = errors.New("auto closed wss conn by server")
)
