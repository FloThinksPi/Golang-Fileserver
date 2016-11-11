// ONLY FOR TESTING RANDOM CODE; CAN BE ALTERED OR DELETED ANY TIME;
package main

import (
	"github.com/pkg/errors"

	"Utils"
)

func main() {
	err := errors.New("Erster Fheler")
	err = errors.Wrap(err, "Zweiter Fehler")
	err = errors.Wrap(err, "Letzter Fehler")

	err2 := cause(err)

	//Utils.HandleErrorPrint(err2)

	Utils.HandlePanic(err2)

	//log.Print(err2)
}

func cause(err error) error {
	err2 := errors.Errorf("Error: %+v", err)
	//err2 := errors.New("ErrorTest")
	return err2

}