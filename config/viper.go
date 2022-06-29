package config

import (
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
)

var ViperConfig Configuration

func init() {
	runtimeViper := viper.New()
	// 设置配置文件
	runtimeViper.AddConfigPath(".")
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("json")
	// 读取所有配置
	err := runtimeViper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 将读取到的配置文件映射到 Configuration 结构体变量 ViperConfig
	runtimeViper.Unmarshal(&ViperConfig)

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active.en.json")
	bundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	ViperConfig.LocaleBundle = bundle

	// 监听配置文件变更，该监听会开启新的协程执行，不影响和阻塞当前协程
	runtimeViper.WatchConfig()
	// 当配置文件有变更时，调用匿名回调函数
	runtimeViper.OnConfigChange(func(e fsnotify.Event) {
		// 重新加载配置文件，并将配置值映射到 ViperConfig 指针
		runtimeViper.Unmarshal(&ViperConfig)
		// 重新加载新的语言文件
		ViperConfig.LocaleBundle.MustLoadMessageFile(ViperConfig.App.Locale + "/active." + ViperConfig.App.Language + ".json")
	})
}