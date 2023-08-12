package conf

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

var (
	Conf *viper.Viper // 全局配置
)

func init() {
	var err error
	Conf, err = InitConfig()
	if err != nil {
		panic(err)
	}
}

func InitConfig() (*viper.Viper, error) {
	v := viper.New()
	v.BindEnv("GO_ENV")
	if v.GetString("GO_ENV") == "prod" {
		v.SetConfigName("prod")
	} else if v.GetString("GO_ENV") == "test" {
		v.SetConfigName("test")
	} else {
		// 当环境变量GO_ENV未设置时，默认使用dev配置
		v.SetConfigName("dev")
	}
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err == nil {
		klog.Info("Using config: ", v.ConfigFileUsed())
	} else {
		return nil, err
	}
	return v, nil
}

func GetConf() *viper.Viper {
	return Conf
}
