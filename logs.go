package logs

import (
	"log"
	"os"
)

/*
@Time : 2020/11/30 11:52 下午
@Author : zhuyongsheng
*/

const (
	ERROR int = iota
	WARN
	INFO
	DEBUG
)

var (
	error_ *log.Logger
	warn   *log.Logger
	info   *log.Logger
	debug  *log.Logger
	Level  int
)

func InitLogFile(fileName string, level int) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}
	error_ = log.New(file,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	warn = log.New(file,
		"WARN: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	info = log.New(file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	debug = log.New(file,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile)
	Level = level
}

func InitLogDirectory(path string, level int) {
	errorFile, err := os.OpenFile(path+"error_.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error_ log file:", err)
	}
	error_ = log.New(errorFile,
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
	warnFile, err := os.OpenFile(path+"warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open warn log file:", err)
	}
	warn = log.New(warnFile,
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
	infoFile, err := os.OpenFile(path+"info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open info log file:", err)
	}
	info = log.New(infoFile,
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
	debugFile, err := os.OpenFile(path+"debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open debug log file:", err)
	}
	debug = log.New(debugFile,
		"",
		log.Ldate|log.Ltime|log.Lshortfile)
	Level = level
}

func Error(format string, v ...interface{}) {
	if Level >= ERROR {
		error_.Printf(format, v...)
	}
}

func Warn(format string, v ...interface{}) {
	if Level >= WARN {
		warn.Printf(format, v...)
	}
}

func Info(format string, v ...interface{}) {
	if Level >= INFO {
		info.Printf(format, v...)
	}
}

func Debug(format string, v ...interface{}) {
	if Level >= DEBUG {
		debug.Printf(format, v...)
	}
}

// ignore the level set
func Log(level int, format string, v ...interface{}) {
	switch level {
	case ERROR:
		error_.Printf(format, v...)
	case WARN:
		warn.Printf(format, v...)
	case INFO:
		info.Printf(format, v...)
	case DEBUG:
		debug.Printf(format, v...)
	}
}
