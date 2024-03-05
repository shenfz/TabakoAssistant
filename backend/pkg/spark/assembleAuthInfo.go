package spark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/shenfz/TabakoAssistant/backend/config"
	"net/url"
	"strings"
	"time"
)

/**
 * @Author shenfz
 * @Date 2024/3/4 9:57
 * @Email 1328919715@qq.com
 * @Description:
 **/

// AssembleAuthUrl 创建鉴权url  apikey 即 hmac username
func AssembleAuthUrl() (string, error) {
	sparkConf := config.GetAppConfig()
	ul, err := url.Parse(sparkConf.Spark.HostUrl)
	if err != nil {
		return "", err
	}
	sha := hMacWithShaToBase64("hmac-sha256", getSign(ul), sparkConf.Spark.AppSecret)
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", sparkConf.Spark.AppKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", time.Now().UTC().Format(time.RFC1123))
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callUrl := sparkConf.Spark.HostUrl + "?" + v.Encode()
	return callUrl, nil
}

func getSign(ul *url.URL) string {
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	return strings.Join(signString, "\n")
}

func hMacWithShaToBase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}
