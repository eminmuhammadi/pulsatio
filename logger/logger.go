package logger

import (
	"log"
	"os"
)

func Debug() bool {
	return os.Getenv("PULSATIO_DEBUG") != "false"
}

func Print(msg string) {
	if Debug() {
		log.Println("[pulsatio] " + msg)
	}
}

func Printf(format string, v ...interface{}) {
	if Debug() {
		log.Printf("[pulsatio] "+format+"\n", v...)
	}
}

func Fatal(msg string) {
	if Debug() {
		log.Fatalln("[pulsatio] " + msg)
	}
}

func Fatalf(format string, v ...interface{}) {
	if Debug() {
		log.Fatalf("[pulsatio] "+format+"\n", v...)
	}
}
