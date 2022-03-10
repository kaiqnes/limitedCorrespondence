package main

import (
	"fmt"
	"time"
)

func executeFuncWithTimeTrack(name string, function func()) {
	init := time.Now()
	function()
	fmt.Printf("Func %s took %s\n", name, time.Since(init).String())
}
