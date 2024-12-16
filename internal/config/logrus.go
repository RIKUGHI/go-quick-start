package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

var LogsDirectory = "./logs"

// type Logger struct {
// 	*logrus.Logger
// 	file        *os.File
// 	currentDate string
// }

func NewLogger() *logrus.Logger {
	log := logrus.New()

	log.SetLevel(logrus.Level(6))
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.DateTime,
	})

	// logger := &Logger{
	// 	Logger: log,
	// }

	// logger.setLogFile()

	return log
}

// func (l *Logger) setLogFile() {
// 	currentDate := time.Now().Format(time.DateOnly)
// 	if l.currentDate != currentDate {
// 		if l.file != nil {
// 			l.file.Close()
// 		}

// 		logDir := fmt.Sprintf("%s/%s/", LogsDirectory, strings.ReplaceAll(currentDate, "-", "/"))

// 		if _, err := os.Stat(logDir); os.IsNotExist(err) {
// 			if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
// 				log.Printf("error creating logs directory: %v", err)
// 				return
// 			}
// 		}

// 		logFileName := fmt.Sprintf("%s%s.log", logDir, currentDate)
// 		logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

// 		if err != nil {
// 			l.Logger.Fatalf("Error opening log file: %v", err)
// 		}

// 		l.Logger.SetOutput(logFile)

// 		l.file = logFile
// 		l.currentDate = currentDate
// 	}
// }

// func (l *Logger) Info(args ...interface{}) {
// 	l.setLogFile()

// 	l.Logger.Info(args...)
// }

// func (l *Logger) Printf(format string, args ...interface{}) {
// 	l.setLogFile()

// 	l.Logger.Printf(format, args)
// }

// func (l *Logger) Fatalf(format string, args ...interface{}) {
// 	l.setLogFile()

// 	l.Logger.Logf(logrus.FatalLevel, format, args...)
// }

// func (l *Logger) Errorf(format string, args ...interface{}) {
// 	l.Logger.Logf(logrus.ErrorLevel, format, args...)
// }

// func (l *Logger) Println(args ...interface{}) {
// 	l.setLogFile()

// 	l.Logger.Println(args)
// }
