package config

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// 数据库配置项
var cfgDatabase *viper.Viper

// 应用配置项
var cfgAppliction *viper.Viper

// Token配置项
var cfgJwt *viper.Viper

// Log配置项
var cfgLogger *viper.Viper

// Ssl配置项 非必须
var cfgSsl *viper.Viper

// 代码生成配置项 非必须
var cfgGen *viper.Viper

func Setup(path string) {
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}
	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("No found settings.database in the configuration")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgAppliction =  viper.Sub("settings.appliction")
	if cfgAppliction == nil {
		panic("No found settings.appliction in the configuration")
	}
	ApplictionConfig = InitAppliction(cfgAppliction)

	cfgJwt =  viper.Sub("settings.Jwt")
	if cfgJwt == nil {
		panic("No found settings.Jwt in the configuration")
	}
	JwtConfig = InitJwt(cfgJwt)

	cfgLogger =  viper.Sub("settings.logger")
	if cfgLogger == nil {
		panic("No found settings.logger in the configuration")
	}
	LoggerConfig = InitLogger(cfgLogger)

	cfgSsl = viper.Sub("settings.ssl")
	if cfgSsl == nil {
		// Ssl不是系统强制要求的配置，默认可以不用配置，将设置为关闭状态
		fmt.Println("warning config not found settings.ssl in the configuration")
		SslConfig = new(Ssl)
		SslConfig.Enable = false
	} else {
		SslConfig = InitSsl(cfgSsl)
	}

	cfgGen = viper.Sub("setting.gen")
	if cfgGen == nil {
		panic("No found settings.gen")
	}
	GenConfig = InitGen(cfgGen)

}
