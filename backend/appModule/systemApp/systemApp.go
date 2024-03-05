package systemApp

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/getlantern/systray"
	"github.com/shenfz/TabakoAssistant/backend/global"
	"github.com/skratchdot/open-golang/open"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

/**
 * @Author shenfz
 * @Date 2024/1/20 19:03
 * @Email 1328919715@qq.com
 * @Description:
 **/

//go:embed icon.ico
var iconsEmbed []byte

// SystemApp struct
type SystemApp struct {
	ctx          context.Context
	enableSpeech bool
	callbacks    []func()
}

// NewSystemApp creates a new SystemApp SystemApplication struct
func NewSystemApp() *SystemApp {
	return &SystemApp{
		ctx:          context.Background(),
		enableSpeech: true,
		callbacks:    make([]func(), 0),
	}
}

func (a *SystemApp) Startup(ctx context.Context) {
	a.ctx = ctx
	// 建立监听事件，无数据处理
	a.callbacks = append(a.callbacks, runtime.EventsOn(a.ctx, global.EventToggleSpeechAPI, func(optionalData ...interface{}) {
		fmt.Printf("====> %v \n", optionalData)
	}))
	systray.Run(a.runReady, a.runExit)
}

func (a *SystemApp) runReady() {
	systray.SetIcon(iconsEmbed)
	systray.SetTitle(global.AppName)
	systray.SetTooltip(global.AppTips)
	mOpenMain := systray.AddMenuItem("主页", "open the main window")
	systray.AddSeparator()
	mReference := systray.AddMenuItem("技术栈", "wails,vue3,tailwind,element-plus")
	mRefSub1 := mReference.AddSubMenuItem("wails", "https://wails.io/zh-Hans/docs")
	mRefSub2 := mReference.AddSubMenuItem("vue3", "https://cn.vuejs.org/")
	mRefSub3 := mReference.AddSubMenuItem("tailwind", "https://www.tailwindcss.cn/docs/installation")
	mRefSub4 := mReference.AddSubMenuItem("element-plus", "https://element-plus.org/zh-CN")
	systray.AddSeparator()
	mSpeechAPIEnable := systray.AddMenuItemCheckbox("开启语音识别", "借助浏览器的API实现", true)
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mOpenMain.ClickedCh:
				runtime.WindowShow(a.ctx)
			case <-mRefSub1.ClickedCh:
				open.Run("https://wails.io/zh-Hans/docs")
			case <-mRefSub2.ClickedCh:
				open.Run("https://cn.vuejs.org/")
			case <-mRefSub3.ClickedCh:
				open.Run("https://www.tailwindcss.cn/docs/installation")
			case <-mRefSub4.ClickedCh:
				open.Run("https://element-plus.org/zh-CN")
			case <-mSpeechAPIEnable.ClickedCh:
				if mSpeechAPIEnable.Checked() {
					mSpeechAPIEnable.Uncheck()
				} else {
					mSpeechAPIEnable.Check()
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
	// We can manipulate the systray in other goroutines
	// go func() {
	//mChange := systray.AddMenuItem("Change Me", "Change Me")
	//// mAllowRemoval := systray.AddMenuItem("Allow removal", "macOS only: allow removal of the icon when cmd is pressed")
	//mChecked := systray.AddMenuItemCheckbox("Unchecked", "Check Me", true)
	//mEnabled := systray.AddMenuItem("Enabled", "Enabled")
	//
	//subMenuTop := systray.AddMenuItem("SubMenuTop", "SubMenu Test (top)")
	//subMenuMiddle := subMenuTop.AddSubMenuItem("SubMenuMiddle", "SubMenu Test (middle)")
	//subMenuBottom := subMenuMiddle.AddSubMenuItemCheckbox("SubMenuBottom - Toggle Panic!", "SubMenu Test (bottom) - Hide/Show Panic!", false)
	//subMenuBottom2 := subMenuMiddle.AddSubMenuItem("SubMenuBottom - Panic!", "SubMenu Test (bottom)")

	// mUrl := systray.AddMenuItem("Open UI", "my home")\

}

func (a *SystemApp) runExit() {
	for _, callback := range a.callbacks {
		callback()
	}
	runtime.Quit(a.ctx)
}

func (a *SystemApp) ToggleSpeechAPI() {
	a.enableSpeech = !a.enableSpeech
}

func (a *SystemApp) GetSpeechAPIEventName() string {
	return global.EventToggleSpeechAPI
}
