package Utils

import (
	"log"
	"github.com/pkg/errors"
)

func HandleErrorPrint(err error) {
	if err != nil {
		LogError(errors.Cause(err).Error())
	}
}

func HandlePanic(err error) {
	if err != nil {
		log.Panic("!!!PANIC!!! : ", errors.Cause(err))
	}
}

