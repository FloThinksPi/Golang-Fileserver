package Flags

import "flag"

var Verbosity = flag.Int("v",3,"Verbosity (0=Errors ,3=All)")

func init() {
	flag.Parse()
}
