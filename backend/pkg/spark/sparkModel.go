package spark

import (
	"github.com/shenfz/TabakoAssistant/backend/config"
	"github.com/shenfz/TabakoAssistant/backend/pkg/snowFlakeID"
)

/**
 * @Author shenfz
 * @Date 2024/2/4 14:45
 * @Email 1328919715@qq.com
 * @Description:
 **/

const (
	Spark_With_Context = iota
	Spark_Not_With_Context
)

type ParamsOptions interface {
	Apply(*SparkReqMsg)
}

type ReqHeader struct {
	AppId string `json:"app_id"`
	Uid   string `json:"uid"` // 	每个用户的id，用于区分不同用户
}

type ReqParameter struct {
	Chat ReqChat `json:"chat"`
}

type ReqChat struct {
	Domain      string  `json:"domain"`      // generalv3.5指向V3.5版本;
	Temperature float64 `json:"temperature"` // 取值范围 (0，1] ，默认值0.5 核采样阈值。决定结果随机性，取值越高随机性越强即相同的问题得到的不同答案的可能性越高
	MaxTokens   int     `json:"max_tokens"`  // V3.5取值为[1,8192]  模型回答的tokens的最大长度
	ChatId      string  `json:"chat_id"`     // 唯一ID 用于关联用户会话
}

type ReqPayload struct {
	Message ReqMessage `json:"message"`
}

type ReqMessage struct {
	Text []Text `json:"text"`
}

type Text struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ============================================>
type SparkReqMsg struct {
	ReqHeader    `json:"header"`
	ReqParameter `json:"parameter"`
	ReqPayload   `json:"payload"`
}

func NewSparkRequestMsg(opts ...ParamsOptions) SparkReqMsg {
	s := config.GetAppConfig().Spark
	obj := SparkReqMsg{
		ReqHeader: ReqHeader{AppId: s.AppID, Uid: snowFlakeID.GetSnowFlakeID()},
		ReqParameter: ReqParameter{ReqChat{
			Domain:      s.SparkApiDomain,
			Temperature: 0.6,
			MaxTokens:   8192,
			ChatId:      "Chat" + snowFlakeID.GetSnowFlakeID(),
		}},
		ReqPayload: ReqPayload{
			Message: ReqMessage{Text: []Text{}},
		},
	}
	for _, opt := range opts {
		opt.Apply(&obj)
	}
	return obj
}

// ==========================================>
type SparkRespMsg struct {
	RespHeader  `json:"header"`
	RespPayload `json:"payload"`
}

type RespHeader struct {
	Code    int    `json:"code"` // ，0表示正常，非0表示出错
	Message string `json:"message"`
	Sid     string `json:"sid"`    // 会话的唯一id，用于讯飞技术人员查询服务端会话日志使用,出现调用错误时建议留存该字段
	Status  int    `json:"status"` // 0代表首次结果；1代表中间结果；2代表最后一个结果
}

type RespPayload struct {
	RespChoices `json:"choices"`
	RespUsage   `json:"usage"`
}

type RespUsage struct {
	RespUsageText `json:"text"`
}

type RespUsageText struct {
	QuestionTokens   int `json:"question_tokens"`   // 保留字段，可忽略
	PromptTokens     int `json:"prompt_tokens"`     // 包含历史问题的总tokens大小
	CompletionTokens int `json:"completion_tokens"` // 回答的tokens大小
	TotalTokens      int `json:"total_tokens"`      // rompt_tokens和completion_tokens的和，也是本次交互计费的tokens大小
}

type RespChoices struct {
	Status int    `json:"status"` // 文本响应状态，取值为[0,1,2]; 0代表首个文本结果；1代表中间文本结果；2代表最后一个文本结果
	Seq    int    `json:"seq"`    // 数据序号 [0,9999999]
	Text   []Text `json:"text"`   //   RespChoicesText
}

type RespChoicesText struct {
	Content string `json:"content"` // 回答内容
	Role    string `json:"role"`    // 	角色标识，固定为assistant，标识角色为AI
	Index   int    `json:"index"`   // 保留字段，开发者可忽略
}
