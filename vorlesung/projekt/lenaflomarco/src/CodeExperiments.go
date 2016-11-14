// ONLY FOR TESTING RANDOM CODE; CAN BE ALTERED OR DELETED ANY TIME;
package main

import (
	"Utils"
	"time"
	"github.com/pkg/errors"
)

func main() {
	Utils.LogError("Error")
	Utils.LogWarning("Test")
	Utils.LogInfo("TEst")
	Utils.LogDebug("TAda")

	time.Sleep(time.Second * 1)

	err := errors.New("TEST")
	Utils.HandlePrint(err)

	time.Sleep(time.Second * 1)
	Utils.HandlePanic(err)
}

