package main

import "testing"

//maintest - template unittest
func maintest(t *testing.T) {
	if testfunction(1) != 1 {
		t.Error("Test Failed")
	}
}
