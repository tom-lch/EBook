package api

import (
	"admin/database"
	"admin/global"
	"admin/models/gorm"
	"admin/pkg/logger"
	"admin/tools/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	ConfigYaml string
	port string
	mode string
	StartCmd = &cobra.Command{
		Use: "server",
		Short: "Start API server",
		Example: "./admin server -c config/settings.yaml",
		SilenceUsage: true,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	echoTimes int
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&ConfigYaml, "config", "-c", "config/settings", "start server with provied configuration file")
	StartCmd.PersistentFlags().StringVarP(&port, "port", "-p", ":8000", "TCP port server listering on")
	StartCmd.PersistentFlags().StringVarP(&mode, "mode", "-m", "dev", "server mode; eg: dev, test, pord")
}

func setup() {
	fmt.Println("对服务进行初始化。。。。")
	// 1.读取配置文件 将配置信息读入变量中√
	config.Setup(ConfigYaml)
	// 2. 设置日志 将日志配置到全局变量中√
	logger.Setup()
	// 3. 初始化数据库连接 将数据库连接配置到全局变量中
	database.Setup(config.DatabaseConfig.Driver)
	// 4. 接口访问控制
	// casbin.Setup()
	// 数据库迁移
	_ = MigrateModel()
	usageStr := `starting api server`
	global.Logger.Info(usageStr)

}

func MigrateModel() error {
	if config.DatabaseConfig.Driver == "mysql" {
		global.GDB = global.GDB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	}
	return gorm.AutoMigrate(global.GDB)
}

func run() error {
	// 根据读取的日志信息启动web
	if viper.GetString("settings.application.mode") == string(tools.ModeProd) {
		gin.SetMode(gin.ReleaseMode)
	}

}