package global

import "fmt"

var Value int

func init() {
	fmt.Println("global is being imported")
	Value = 3
}
