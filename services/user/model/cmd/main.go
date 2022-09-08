package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"mxshop/services/user/model"
	"os"
	"time"

	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	dsn := "root:123456@tcp(39.98.125.151:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 彩色打印
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		log.Panicln(err)
	}

	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("li%d", i),
			Mobile:   fmt.Sprintf("1871072862%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}

	db.AutoMigrate(&model.User{})

	// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("generic password", options)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	//check := password.Verify("generic password", salt, encodedPwd, options)
	//fmt.Println(check) // true
	//fmt.Println(genMd5("123456"))
}

func genMd5(code string) string {
	MD5 := md5.New()
	io.WriteString(MD5, code)
	return hex.EncodeToString(MD5.Sum(nil))
}
