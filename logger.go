package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// AppName, Level, Message, Json Object(s), Sourcecode Info
const logFormat = "[%s][%s] %s %s, (%s)"

// Fatal accepts message plus objects or use as normal printf function
func Fatal(msg string, obj ...interface{}) {
	message, jsonObject := checkForFormatting(msg, obj)
	sInfo := sourceCodeInfo()

	log.Fatal(fmt.Sprintf(logFormat, os.Getenv("APP"), "FATAL", message, jsonObject, sInfo))
}

// Error accepts message plus objects or use as normal printf function
func Error(msg string, obj ...interface{}) {
	message, jsonObject := checkForFormatting(msg, obj)
	sInfo := sourceCodeInfo()

	log.Println(fmt.Sprintf(logFormat, os.Getenv("APP"), "ERROR", message, jsonObject, sInfo))
}

// Warn accepts message plus objects or use as normal printf function
func Warn(msg string, obj ...interface{}) {
	message, jsonObject := checkForFormatting(msg, obj)
	sInfo := sourceCodeInfo()

	log.Println(fmt.Sprintf(logFormat, os.Getenv("APP"), "WARN", message, jsonObject, sInfo))
}

// Info accepts message plus objects or use as normal printf function
func Info(msg string, obj ...interface{}) {
	message, jsonObject := checkForFormatting(msg, obj)
	sInfo := sourceCodeInfo()

	log.Println(fmt.Sprintf(logFormat, os.Getenv("APP"), "INFO", message, jsonObject, sInfo))
}

// Debug accepts message plus objects or use as normal printf function
func Debug(msg string, obj ...interface{}) {
	message, jsonObject := checkForFormatting(msg, obj)
	sInfo := sourceCodeInfo()

	log.Println(fmt.Sprintf(logFormat, os.Getenv("APP"), "DEBUG", message, jsonObject, sInfo))
}

func checkForFormatting(msg string, obj []interface{}) (message string, jsonObject []byte) {
	// Fatal accepts message plus objects or use as normal printf function
	// If the msg variable contains a % sign it is considered a printf so we treat it as such
	// Else we will just marshal the object for a separate message and json object
	if strings.Contains(msg, "%") && len(obj) > 0 || obj == nil {
		message = fmt.Sprintf(msg, obj...)
	} else {
		message = msg
		jsonObject, _ = json.Marshal(obj)
	}
	return
}

func sourceCodeInfo() string {
	pc, file, line, _ := runtime.Caller(2)
	funcCall := runtime.FuncForPC(pc).Name()
	return fmt.Sprintf("FILE:%s, FUNC:%s, LINE:%s", file, funcCall, strconv.Itoa(line))
}
