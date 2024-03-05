package test

import (
	"context"
	"github.com/shenfz/TabakoAssistant/backend/config"
	"github.com/shenfz/TabakoAssistant/backend/pkg/spark"
	"testing"
	"time"
)

/**
 *  WebAPI 接口调用示例 接口文档（必看）：https://www.xfyun.cn/doc/spark/Web.html
 * 错误码链接：https://www.xfyun.cn/doc/spark/%E6%8E%A5%E5%8F%A3%E8%AF%B4%E6%98%8E.html（code返回错误码时必看）
 * @author shenfz
 */

func Test_LoadJson(t *testing.T) {
	config.GetAppConfig()
}

func Test_AssembleURL(t *testing.T) {
	config.GetAppConfig()
	url, err := spark.AssembleAuthUrl()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(url)
}

func Test_GenParams(t *testing.T) {
	config.GetAppConfig()
	t.Logf("%+v", spark.NewSparkRequestMsg())
}

func Test_Simple(t *testing.T) {
	config.GetAppConfig()
	s := spark.GetSparkWsConn(context.Background())

	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		s.TestEcho()
	}
	select {}
}
