package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"qxshop_srvs/user_srv/model"
	"time"

	"github.com/anaskhan96/go-password-encoder"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

func main() {
	dsn := "root:123456@tcp(192.168.43.166:3306)/qxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)
	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "qx_",
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("test123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("qingxin%d", i),
			Mobile:   fmt.Sprintf("131456783%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}

	//_ = db.AutoMigrate(&model.User{})
	// Using the default options
	/*	salt, encodedPwd := password.Encode("generic password", nil)
		check := password.Verify("generic password", salt, encodedPwd, nil)
		fmt.Println(check) // true
	*/

	// Using custom options
	/*	options := &password.Options{16, 100, 32, sha512.New}
		salt, encodedPwd := password.Encode("generic password", options)
		newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
		fmt.Println(len(newPassword))
		fmt.Println(newPassword)

		passwordInfo := strings.Split(newPassword, "$")
		fmt.Println(passwordInfo)
		check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
		fmt.Println(check) // true*/
}
