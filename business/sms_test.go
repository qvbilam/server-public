package business

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"public/global"
	"testing"
	"time"
)

func newLogger() logger.Interface {
	log.Writer()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,       // 禁用彩色打印
		},
	)
	return newLogger
}

func initCache() {
	host := "127.0.0.1"
	port := 6379
	global.Redis = *redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", host, port),
		DB:   0,
	})
}
func initDB() {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	database := "qvbilam_file"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
		Logger: newLogger(),
	})
	if err != nil {
		panic(any(err))
	}
	global.DB = db
}

func TestSmsBusiness_Send(t *testing.T) {
	initCache()
	initDB()
	b := SmsBusiness{
		Mobile:   "13501294164",
		Type:     "login",
		ClientIP: "127.0.0.1",
	}
	err := b.Send()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestSmsBusiness_Check(t *testing.T) {
	initCache()
	initDB()
	b := SmsBusiness{
		Mobile:    "15032061937",
		Type:      "login",
		CheckCode: "7037",
	}
	err := b.Check()
	if err != nil {
		fmt.Println(err.Error())
	}
}
