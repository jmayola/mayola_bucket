package middleware

import (
	"io"
	"log"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var (
	logFile *os.File
	stdout  io.Writer
	once    sync.Once
)

func MiddleLogger() fiber.Handler {
	once.Do(func() {
		err := os.MkdirAll("logs/", 0755)
		if err != nil {
			log.Fatalf("error creating logs folder: %v", err)
		}

		logFile, err = os.OpenFile("logs/api.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		stdout = io.MultiWriter(os.Stdout, logFile)
	})

	loggerConfig := logger.Config{
		Format:     "${time},${ip},${pid},${locals:requestid},${status},${method},${path},${error}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     stdout,
	}

	return logger.New(loggerConfig)
}

// close the file when application shuts down
func CloseLogFile() {
	if logFile != nil {
		logFile.Close()
	}
}
