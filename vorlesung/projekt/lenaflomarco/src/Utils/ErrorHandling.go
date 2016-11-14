package Utils

import (
	"github.com/pkg/errors"
	"time"
	"fmt"
)

func HandlePrint(err error) {
	if err != nil {
		LogErrorWithStack(err)
	}
}

func HandlePanic(err error) {
	if err != nil {
		LogPanic(errors.Cause(err).Error())
		fmt.Print("		")
		time.Sleep(time.Millisecond * 100)// For propper output on console
		panic("Starting Panic Process")
	}
}

