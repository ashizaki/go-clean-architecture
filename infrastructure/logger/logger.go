package logger

import (
	"log"
	"os"
)

// Logger is Log object.
var Logger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "go-clean-architecture", log.LstdFlags)
}
