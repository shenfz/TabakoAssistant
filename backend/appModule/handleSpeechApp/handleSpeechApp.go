package handleSpeechApp

import (
	"context"
	"github.com/shenfz/TabakoAssistant/backend/global"
	"github.com/shenfz/TabakoAssistant/backend/pkg/spark"
	"github.com/shenfz/TabakoAssistant/backend/pkg/zapLogger"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

/**
 * @Author shenfz
 * @Date 2024/2/4 13:25
 * @Email 1328919715@qq.com
 * @Description:
 **/

// HandleSpeechApp struct
type HandleSpeechApp struct {
	ctx       context.Context
	sparkFunc *spark.SparkWsConn
}

// NewHandleSpeechApp creates a new NewHandleSpeechApp HandleSpeech struct
func NewHandleSpeechApp() *HandleSpeechApp {
	return &HandleSpeechApp{}
}

// HandleOrder accept content from speech
func (h *HandleSpeechApp) HandleOrder(inputStr string) {
	//fmt.Println("Get: ", inputStr)
	//return fmt.Sprintf("Time: %v, Echo:%s", time.Now().String(), inputStr)
	//
	h.UseSparkNaturalDialogue(inputStr)
}

func (h *HandleSpeechApp) Startup(ctx context.Context) {
	h.ctx = ctx
	h.sparkFunc = spark.GetSparkWsConn(ctx)
}

// UseSparkNaturalDialogue 使用自然对话
func (h *HandleSpeechApp) UseSparkNaturalDialogue(input string) {

	if !h.sparkFunc.IsOnWorking() {
		zapLogger.GetGlobalLogger().Info(" spark not ready")
		runtime.EventsEmit(h.ctx, "showSpark", " spark not ready")
		return
	}

	// 组装发送
	obj := spark.GenParams(spark.Text{
		Role:    global.SparkAPI_Role_User,
		Content: input,
	})

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))

	zapLogger.GetGlobalLogger().Info(h.sparkFunc.SendMsg(ctx, obj))
}

// UseSparkFunctionCall 使用内置插件 功能调用
func UseSparkFunctionCall(input string, funName string) {

}
