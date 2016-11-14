package Utils

import (
	"fmt"
	"time"
)

func LogPanic(msg string) {
	fmt.Print(ANSI_COLOR_RED, timeNow(), "\x1b[31m", "	PANIC:		", msg, ANSI_COLOR_RESET)
}

func LogError(msg string) {
	fmt.Println(ANSI_COLOR_RED, timeNow(), "	Error:		", msg, ANSI_COLOR_RESET)
}

func LogErrorWithStack(err error)  {
	fmt.Print(" ",ANSI_COLOR_RED,timeNow(), "	Error:		 ")
	fmt.Printf("%+v",err)
	fmt.Println(ANSI_COLOR_RESET)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func LogWarning(msg string) {
	if (VERBOSITY >= 1) {
		fmt.Println(ANSI_COLOR_YELLOW, timeNow(), " Warning:	", msg, ANSI_COLOR_RESET)
	}
}

func LogInfo(msg string) {
	if (VERBOSITY >= 2) {
		fmt.Println(ANSI_COLOR_MAGENTA, timeNow(), " Info:		", msg, ANSI_COLOR_RESET)
	}
}

func LogDebug(msg string) {
	if (VERBOSITY >= 3) {
		fmt.Println(ANSI_COLOR_RESET, timeNow(), " Debug:		", msg, ANSI_COLOR_RESET)
	}
}

func timeNow() string {
	return time.Now().Format(time.StampNano)
}