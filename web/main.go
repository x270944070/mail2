package main

import (
	"fmt"
	"mail.web/config"
	"mail.web/logging"
	"mail.web/router"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need config file.eg: config.yaml")
		return
	}
	// 1.加载配置
	config.Init(os.Args[1])
	// 2.初始化日志
	logging.Init()
	// 3.初始化mysql
	mysql.Init(config.Conf.MySQLConfig)
	// 4.如果参数为 migrate就初始化表结构
	if len(os.Args) >= 3 && os.Args[2] == "migrate" {
		mysql.AutoMigrateDB()
		fmt.Println("run AutoMigrate success!")
		return
	}
	// 5.初始化redis
	redis.Init(setting.Conf.RedisConfig)
	// 6.初始化定时任务
	jobs.InitJobs()
	// 7.注册路由
	r := router.InitRouter()
	r.Run(fmt.Sprintf(":%d", setting.Conf.Port))
}
