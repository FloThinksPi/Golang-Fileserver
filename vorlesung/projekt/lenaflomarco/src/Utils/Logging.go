package Utils

import (
	"fmt"
	"time"
)

func LogError(msg string) {
	fmt.Println(timeNow() + " Error: " + msg)
}

func LogWarning(msg string) {
	if (VERBOSITY >= 1) {
		fmt.Println(timeNow() + " Warning: " + msg)
	}
}

func LogInfo(msg string) {
	if (VERBOSITY >= 2) {
		fmt.Println(timeNow() + " Info: " + msg)
	}
}

func LogDebug(msg string) {
	if (VERBOSITY >= 3) {
		fmt.Println(timeNow() + " Debug: " + msg)
	}
}

func timeNow() string {
	return time.Now().Format(time.RFC1123)
}