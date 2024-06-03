package conf

import (
	"gin_mall_tmp/dao"
	"strings"

	"gopkg.in/ini.v1"
)

var (
	AppModel string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	ValidEmail string
	SmtpHost   string
	SmtpEmail  string
	SmtpPass   string
	SmtpPort   int

	Host        string
	ProductPath string
	AvatarPath  string
)

// 本地讀取環境變量
func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}
	LoadServer(file)
	LoadMySql(file)
	LoadRedis(file)
	LoadEmail(file)
	LoadPhotoPath(file)

	// mysql 讀(8成)
	pathRead := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	// mysql 寫(2成)
	pathWrite := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	dao.Database(pathRead, pathWrite) // 初始化数据库
}

func LoadServer(file *ini.File) {
	AppModel = file.Section("service").Key("AppMode").MustString("debug")
	HttpPort = file.Section("service").Key("HttpPort").MustString(":8080")
}

func LoadMySql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").MustString("mysql")
	DbHost = file.Section("mysql").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("mysql").Key("DbPort").MustString("3307")
	DbUser = file.Section("mysql").Key("DbUser").MustString("root")
	DbPassWord = file.Section("mysql").Key("DbPassWord").MustString("11111111")
	DbName = file.Section("mysql").Key("DbName").MustString("mysql")
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").MustString("0")
	RedisAddr = file.Section("redis").Key("RedisAddr").MustString("localhost:6379")
	RedisPw = file.Section("redis").Key("RedisPw").MustString("")
	RedisDbName = file.Section("redis").Key("RedisDbName").MustString("2")
}

func LoadEmail(file *ini.File) {
	ValidEmail = file.Section("email").Key("ValidEmail").MustString("http://localhost:8080/#/valid/email/")
	SmtpHost = file.Section("email").Key("SmtpHost").MustString("smtp.gmail.com")
	SmtpEmail = file.Section("email").Key("SmtpEmail").MustString("")
	SmtpPass = file.Section("email").Key("SmtpPass").MustString("")
	SmtpPort = file.Section("email").Key("SmtpPort").MustInt(587)
}

func LoadPhotoPath(file *ini.File) {
	Host = file.Section("path").Key("Host").MustString("http://127.0.0.1")
	ProductPath = file.Section("path").Key("ProductPath").MustString("/static/imgs/product/")
	AvatarPath = file.Section("path").Key("AvatarPath").MustString("/static/imgs/avatar/")
}
