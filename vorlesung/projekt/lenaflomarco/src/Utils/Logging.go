package Utils

import (
	"fmt"
	"time"
)

func LogError(msg string) {
	fmt.Println(ANSI_COLOR_RED, timeNow(), "\x1b[31m", " Error:	", msg, ANSI_COLOR_RESET)
}

func LogWarning(msg string) {
	if (VERBOSITY >= 1) {
		fmt.Println(ANSI_COLOR_YELLOW, timeNow(), " Warning:	", msg, ANSI_COLOR_RESET)
	}
}

func LogInfo(msg string) {
	if (VERBOSITY >= 2) {
		fmt.Println(ANSI_COLOR_MAGENTA, timeNow(), " Info:	", msg, ANSI_COLOR_RESET)
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