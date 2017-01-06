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
	"github.com/pkg/errors"
	"time"
	"fmt"
)

//HandlePrint handles an error by printig it as Error if err!=nil
func HandlePrint(err error) {
	if err != nil {
		LogErrorWithStack(err)
	}
}

//HandlePanic handels an error by printing it as Panic and causing a panic() if err!=nil
func HandlePanic(err error) {
	if err != nil {
		LogPanic(errors.Cause(err).Error())
		fmt.Print("		")
		time.Sleep(time.Millisecond * 100)// For propper output on console
		panic("Starting Panic Process")
	}
}

