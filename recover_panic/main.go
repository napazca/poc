package main

import (
	"fmt"
	"runtime/debug"

	"github.com/napazca/poc/recover_panic/trigger"
)

func main() {
	// put recover in main function to any catch panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			fmt.Println("Stacktrace from panic: \n" + string(debug.Stack()))
		}
	}()
	child()
}

func child() {
	grandchild()
}

func grandchild() {
	trigger.TPanic()
}
