// Copyright (c) 2019. Shonminh Yang
// db-cacher framework source code and usage is governed by a MIT style license that can be found in the LICENSE file.

package logger

import "fmt"

type LogFunc func(args interface{})

var LogDebug, LogInfo, LogWarn, LogError LogFunc

// register logger function
func RegisterLoggerFunc(logDebugFunc, logInfoFunc, logWarnFunc, logErrorFunc LogFunc) {
	LogDebug = logDebugFunc
	LogInfo = logInfoFunc
	LogWarn = logWarnFunc
	LogError = logErrorFunc
}

func LogDebugf(format string, args ...interface{}) {
	LogDebug(fmt.Sprintf(format, args...))
}

func LogInfof(format string, args ...interface{}) {
	LogInfo(fmt.Sprintf(format, args...))
}

func LogWarnf(format string, args ...interface{}) {
	LogWarn(fmt.Sprintf(format, args...))
}

func LogErrorf(format string, args ...interface{}) {
	LogError(fmt.Sprintf(format, args...))
}
