package configs

import (
	"flag"
	"github.com/bytedance/sonic"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"os"
)

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}
	viper.SetConfigType("yaml")
}

func LoadConfigs() {
	refreshProperties()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		slog.Error("[Properties]无法监控文件事件.", "error=", err.Error())
		return
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					slog.Info("[Properties]检测到文件发生变化")
					refreshProperties()
				}
			case e := <-watcher.Errors:
				if e != nil {
					slog.Info("[Properties]监控文件出错", "error", e.Error())
				}
			}
		}
	}()
	if err = watcher.Add("conf/application.yaml"); err != nil {
		slog.Info("[Properties]添加文件监控失败", "error", err.Error())
	}
}

func refreshProperties() {
	file, err := os.Open("conf/application.yaml")
	if err != nil {
		slog.Error("[Properties]无法打开文件conf/application.yaml", "error", err.Error())
		return
	}
	if err = viper.MergeConfig(file); err != nil {
		slog.Error("[Properties]无法加载配置", "error", err.Error())
	}
	settings, _ := sonic.MarshalString(viper.AllSettings())
	slog.Info("[Properties]已刷新配置.", "settings", settings)
}
