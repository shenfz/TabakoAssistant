package global

/**
 * @Author shenfz
 * @Date 2024/1/20 18:20
 * @Email 1328919715@qq.com
 * @Description:
 **/

const (
	AppName = "TabakoAssistant"
	AppTips = "一个练习做的小工具罢了"
)

const (
	EventToggleSpeechAPI = "Toggle_Speech_API" // 切换语音识别开关
)

const (
	APPID      = "5e7d6ae2"
	APP_Secret = "9e3e909c94e2554964f3d80c0c2904be"
	APP_KEY    = "0ddcdbd0dce5263eda501dde653aa611"

	UUID                = ""
	SparkAPI_Domain     = "generalv3.5"
	SparkAPI_Max_Tokens = 8192 // 最大回答长度
)

const (
	SparkAPI_3_Point_5_URL  = "wss://spark-api.xf-yun.com/v3.5/chat"
	SparkAPI_Role_User      = "user" // system用于设置对话背景，user表示是用户的问题，assistant表示AI的回复
	SparkAPI_Role_System    = "system"
	SparkAPI_Role_Assistant = "assistant"
)
