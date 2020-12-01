package logs

import (
	"log"
	"os"
)

/*
@Time : 2020/11/30 11:52 下午
@Author : zhuyongsheng
*/

var (
	Info  *log.Logger // 重要的信息
	Error *log.Logger // 非常严重的问题
)

func InitLog(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	Info = log.New(file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(file,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}
