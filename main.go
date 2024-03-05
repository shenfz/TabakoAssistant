package main

import (
	"context"
	"embed"
	"github.com/shenfz/TabakoAssistant/backend/appModule/handleSpeechApp"
	"github.com/shenfz/TabakoAssistant/backend/appModule/mainApp"
	"github.com/shenfz/TabakoAssistant/backend/appModule/systemApp"
	"github.com/shenfz/TabakoAssistant/backend/global"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	// wails generate module
	app := mainApp.NewMainApp()
	sysTp := systemApp.NewSystemApp()
	speech := handleSpeechApp.NewHandleSpeechApp()

	//

	// Create application with options
	err := wails.Run(&options.App{
		Title:     global.AppName,
		Width:     1024,
		Height:    300,
		Frameless: true, // 无边框
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent:                true, // 视图透明
			WindowIsTranslucent:                 false,
			DisableWindowIcon:                   false,
			IsZoomControlEnabled:                false,
			ZoomFactor:                          0,
			DisableFramelessWindowDecorations:   false,
			WebviewUserDataPath:                 "",
			WebviewBrowserPath:                  "",
			Theme:                               0,
			CustomTheme:                         nil,
			BackdropType:                        windows.Auto, //
			Messages:                            nil,
			ResizeDebounceMS:                    0,
			OnSuspend:                           nil,
			OnResume:                            nil,
			WebviewGpuIsDisabled:                false,
			WebviewDisableRendererCodeIntegrity: false,
			EnableSwipeGestures:                 false,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) { // 启动器
			app.Startup(ctx)
			sysTp.Startup(ctx)
			speech.Startup(ctx)
		},
		Bind: []interface{}{
			app,
			sysTp,
			speech,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
