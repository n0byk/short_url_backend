package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	os.Exit(0)
}
