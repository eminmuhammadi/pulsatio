package logger

import "log"

func Print(msg string) {
	log.Println("[pulsatio] " + msg)
}

func Printf(format string, v ...interface{}) {
	log.Printf("[pulsatio] "+format+"\n", v...)
}

func Fatal(msg string) {
	log.Fatalln("[pulsatio] " + msg)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf("[pulsatio] "+format+"\n", v...)
}
