package util

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var Logrus *logrus.Logger

func init() {
	src, err := setOutPutFile()
	if err != nil {
		log.Panicln("Failed to set log output file:", err)
	}

	if Logrus != nil {
		Logrus.Out = src // 將日誌輸出到src
		log.Println("Logrus is already initialized, setting output.")
		return
	}
	// 實力化
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel) // 設定日誌級別
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.SetReportCaller(true) // 是否报告调用者信息，在日志中显示调用日志的文件名和行号。

	Logrus = logger
}

// setOutPutFile 日誌輸出
func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	// os.Getwd() 获取当前工作目录路径。
	if dir, err := os.Getwd(); err == nil {
		logFilePath = path.Join(dir, "logs")
	} else {
		return nil, err
	}
	// 检查文件或目录的存在性、权限、大小、修改时间等信息。
	_, err := os.Stat(logFilePath) // 获取指定路径下文件的相关信息。
	if os.IsNotExist(err) {        // 判断文件是否存在
		if err = os.MkdirAll(logFilePath, 0777); err != nil {
			log.Panicln(err.Error())
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	logFileName := now.Format("2006-01-02_15-04-05") + ".log"

	// 日誌文件
	fileName := path.Join(logFilePath, logFileName)
	// 寫入文件
	// O_APPEND : 表示在写入文件时将数据追加到文件末尾，而不是覆盖已有内容。
	// O_CREATE : 如果文件不存在，则创建该文件。
	// O_WRONLY : 表示以只写方式打开文件，即文件只能用于写入数据。
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return src, nil
}
