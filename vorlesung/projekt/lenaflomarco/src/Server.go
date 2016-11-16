package main

import (
	"Utils"
	"Flags"
)

func init() {

}

func main() {

	Utils.LogError("Test")
	Utils.LogWarning("Test")
	Utils.LogInfo("Test")
	Utils.LogDebug(*Flags.WorkDir)

}

