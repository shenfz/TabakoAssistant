package config

import (
	"encoding/json"
	"github.com/shenfz/TabakoAssistant/backend/pkg/zapLogger"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

/**
 * @Author shenfz
 * @Date 2024/1/20 19:05
 * @Email 1328919715@qq.com
 * @Description:
 **/

type AppConfig struct {
	AbsWorkDir string
	Spark      SparksSetting
	logger     *zapLogger.LoggerXC
}

var (
	once                 = sync.Once{}
	appConf              AppConfig
	sparkSettingFileName string = "sparkSetting.json"
)

func GetAppConfig() *AppConfig {
	once.Do(func() {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		separator := string(filepath.Separator)
		appConf.AbsWorkDir = dir[:strings.LastIndex(dir, separator)]
		loadSparkJson()
		appConf.logger = zapLogger.NewLogger()
		appConf.logger.Infof("Get Spark Config: %+v", appConf.Spark)
	})

	return &appConf
}

type SparksSetting struct {
	HostUrl        string `json:"HostUrl"`
	AppID          string `json:"AppID"`
	AppSecret      string `json:"AppSecret"`
	AppKey         string `json:"AppKey"`
	SparkApiDomain string `json:"SparkApiDomain"`
	MaxTokens      int    `json:"MaxTokens"`
}

func loadSparkJson() {
	separator := string(filepath.Separator)
	absFilePath := appConf.AbsWorkDir + separator + "internal" + separator + "config" + separator + sparkSettingFileName
	dataB, err := os.ReadFile(absFilePath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dataB, &appConf.Spark)
	if err != nil {
		panic(err)
	}
}
