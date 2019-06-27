package memusage

import (
	"fmt"
	"io"
	"runtime"
	"time"
)

// MemUsage will write mem profiles to given output file in n intervals
func MemUsage(file io.Writer, n int) {
	go func() {
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
			fmt.Fprintf(file, "\n ===== After ===== %d seconds ", n)
		}
	}()
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
