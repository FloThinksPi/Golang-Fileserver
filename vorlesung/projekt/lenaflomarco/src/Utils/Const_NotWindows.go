/**
  * Fileserver
  * Programmieren II
  *
  * 8376497, Florian Braun
  * 2581381, Lena Hoinkis
  * 9043064, Marco Fuso
 */

// +build !windows

package Utils

// Colors for Terminal , only on Unix systems
const (
	ANSI_COLOR_RED = "\x1b[31m"
	ANSI_COLOR_GREEN = "\x1b[32m"
	ANSI_COLOR_YELLOW = "\x1b[33m"
	ANSI_COLOR_BLUE = "\x1b[34m"
	ANSI_COLOR_MAGENTA = "\x1b[35m"
	ANSI_COLOR_CYAN = "\x1b[36m"
	ANSI_COLOR_RESET = "\x1b[0m"
)