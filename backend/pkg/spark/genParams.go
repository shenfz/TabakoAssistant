package spark

import "github.com/shenfz/TabakoAssistant/backend/global"

/**
 * @Author shenfz
 * @Date 2024/3/4 11:59
 * @Email 1328919715@qq.com
 * @Description:
 **/

type EchoSetOption struct{}

func (EchoSetOption) Apply(s *SparkReqMsg) {
	s.ReqPayload.Message.Text = append([]Text{}, Text{
		Role:    global.SparkAPI_Role_User,
		Content: "Hi!",
	})
	s.ReqParameter.Chat.MaxTokens = 300
	s.ReqParameter.Chat.Temperature = 1
}

// GenParams 生成参数
func GenParams(texts ...Text) SparkReqMsg {
	req := NewSparkRequestMsg()
	req.ReqPayload.Message.Text = append(req.ReqPayload.Message.Text, texts...)
	return req
}
