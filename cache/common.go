package cache

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

var (
	RedisClient *redis.Client
	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("redis load error", err)
	}
	LoadRedis(file)

	Redis()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").MustString("0")
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("localhost:6379")
	RedisPw = file.Section("redis").Key("RedisPw").MustString("")
	RedisDbName = file.Section("redis").Key("RedisDbName").MustString("2")
}

func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)

	RedisClient = redis.NewClient(&redis.Options{
		Addr: RedisAddr,
		// Password: RedisPw,
		DB: int(db),
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
