package memusage

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

// MemUsage will write mem profiles to given output file in n intervals
func MemUsage(filepath string, n int) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	go func() {
		defer file.Close()
		var m runtime.MemStats
		t := time.Now()
		fmt.Fprintf(file, "Time :=> %s\n", t.String())
		for {
			runtime.ReadMemStats(&m)
			// For info on each, see: https://golang.org/pkg/runtime/#MemStats
			fmt.Fprintf(file, "Alloc = %v MiB", bToMb(m.Alloc))
			fmt.Fprintf(file, "\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
			fmt.Fprintf(file, "\tSys = %v MiB", bToMb(m.Sys))
			fmt.Fprintf(file, "\tNumGC = %v\n", m.NumGC)
			time.Sleep(time.Duration(n) * time.Second)
			fmt.Fprintf(file, "\n ==== After ===== %d seconds ", n)
		}
	}()
	return nil
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
