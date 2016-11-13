package Utils

import (
	"github.com/pkg/errors"
)

func HandlePrint(err error) {
	if err != nil {
		LogError(errors.Cause(err).Error())
	}
}

func HandlePanic(err error) {
	if err != nil {
		LogError("!!!PANIC CAUSED!!! : " + errors.Cause(err).Error())
		panic("!! Panicing Now !!!")
	}
}

