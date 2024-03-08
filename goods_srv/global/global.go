package global

import (
	"qxshop_srvs/goods_srv/config"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)

//func init() {
//	dsn := "root:123456@tcp(192.168.1.21:3306)/qx_shop?charset=utf8mb4&parseTime=True&loc=Local"
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold: time.Second, // 慢 SQL 阈值
//			LogLevel:      logger.Info, // Log level
//			Colorful:      false,       // 禁用彩色打印
//		},
//	)
//	// 全局模式
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix:   "qx_",
//			SingularTable: true,
//		},
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(err)
//	}
//}
