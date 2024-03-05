package mainApp

import (
	"context"
	"github.com/shenfz/TabakoAssistant/backend/global"
)

/**
 * @Author shenfz
 * @Date 2024/1/20 18:16
 * @Email 1328919715@qq.com
 * @Description:
 **/

// MainApp struct
type MainApp struct {
	ctx context.Context
}

// NewMainApp creates a new MainApp MainApplication struct
func NewMainApp() *MainApp {
	return &MainApp{}
}

func (a *MainApp) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *MainApp) GetAppName() string {
	return global.AppName
}
