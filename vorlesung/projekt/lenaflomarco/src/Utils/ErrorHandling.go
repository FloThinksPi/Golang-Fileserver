package Utils

import (
	"github.com/pkg/errors"
	"log"
)

func HandlePrint(err error) {
	if err != nil {
		LogError(errors.Cause(err).Error())
	}
}

func HandlePanic(err error) {
	if err != nil {
		log.Panicln("!!!PANIC CAUSED!!! : " + errors.Cause(err).Error())
		defer log.Println("Calling defer Statements now!")
	}
}

