package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Olá, mundo! Estou usando %d processador.\n", numCPU)
}
