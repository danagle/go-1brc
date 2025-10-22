package main

import (
	"os"
	"runtime/pprof"

	v01 "github.com/danagle/go-1brc/v01"
)

func main() {
	// CPU profiling
	cpuFile, _ := os.Create("cpu.prof")
	defer cpuFile.Close()
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()

	inputPath := "measurements.txt"
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	v01.Run(inputPath)
}
