package common

import (
	"runtime"
	"fmt"
)

func PrintMemoryStats() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	fmt.Println(fmt.Sprintf("### Stats: Mem Alloc %d KB, Heap Alloc %d KB", mem.Alloc/1000, mem.HeapAlloc/1000))
}
