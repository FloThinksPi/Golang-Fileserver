package Utils

import (
	"fmt"
	"time"
)

const (
	ANSI_COLOR_RED = "\x1b[31m"
	ANSI_COLOR_GREEN = "\x1b[32m"
	ANSI_COLOR_YELLOW = "\x1b[33m"
	ANSI_COLOR_BLUE = "\x1b[34m"
	ANSI_COLOR_MAGENTA = "\x1b[35m"
	ANSI_COLOR_CYAN = "\x1b[36m"
	ANSI_COLOR_RESET = "\x1b[0m"
)

func LogError(msg string) {
	fmt.Println(timeNow(), "\x1b[31m", " Error:	", msg)
}

func LogWarning(msg string) {
	if (VERBOSITY >= 1) {
		fmt.Println(timeNow(), " Warning:	", msg)
	}
}

func LogInfo(msg string) {
	if (VERBOSITY >= 2) {
		fmt.Println(timeNow(), " Info:	", msg)
	}
}

func LogDebug(msg string) {
	if (VERBOSITY >= 3) {
		fmt.Println(timeNow(), " Debug:	", msg)
	}
}

func timeNow() string {
	return time.Now().Format(time.StampNano)
}