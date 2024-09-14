package main

import (
	"fmt"

	"sandbox/pkg/a"
	"sandbox/pkg/global"
)

func main() {

	fmt.Printf("a.Value: %d\n", a.Value())
	global.Value = 5

	fmt.Printf("a.Value: %d\n", a.Value())
}
