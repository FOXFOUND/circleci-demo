package main

import (
	"fmt"
	"runtime"
)

func main() {
  fmt.Printf("[%v] Golang app (standard)\n",runtime.Version())
}
