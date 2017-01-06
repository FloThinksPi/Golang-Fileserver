/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

package Utils

import (
	"fmt"
	"time"
	"Flags"
)

var verbosity = Flags.GetVerbosity()

//LogPanic Prints a Message as Panic
func LogPanic(msg string) {
	fmt.Print(ANSI_COLOR_RED, timeNow(), "\x1b[31m", "	PANIC:		", msg, ANSI_COLOR_RESET)
}

//LogError Prints a Message as Error
func LogError(msg string) {
	fmt.Println(ANSI_COLOR_RED, timeNow(), "	Error:		", msg, ANSI_COLOR_RESET)
}

//LogErrorWithStack Prints a Error and prints its Stacktrace
func LogErrorWithStack(err error) {
	fmt.Print(" ", ANSI_COLOR_RED, timeNow(), "	Error:		 ")
	fmt.Printf("%+v", err)
	fmt.Println(ANSI_COLOR_RESET)
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

//LogWarning Prints a Message as Warning
func LogWarning(msg string) {
	if (verbosity >= 1) {
		fmt.Println(ANSI_COLOR_YELLOW, timeNow(), " Warning:	", msg, ANSI_COLOR_RESET)
	}
}

//LogInfo Prints a Message as Info
func LogInfo(msg string) {
	if (verbosity >= 2) {
		fmt.Println(ANSI_COLOR_MAGENTA, timeNow(), " Info:		", msg, ANSI_COLOR_RESET)
	}
}

//LogDebug Prints a Message as Debug
func LogDebug(msg string) {
	if (verbosity >= 3) {
		fmt.Println(ANSI_COLOR_RESET, timeNow(), " Debug:		", msg, ANSI_COLOR_RESET)
	}
}

//timeNow outputs a time as string, used as abstraction to make changes easy
func timeNow() string {
	return time.Now().Format(time.StampNano)
}