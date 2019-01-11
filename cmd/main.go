package main

import (
	"got-running-256/internal/iterate"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {

	expo, _ := strconv.Atoi(os.Args[1])

	for goCount := 0; goCount < 255; goCount++ {
		wg.Add(1)
		go func(goCount int) {
			defer wg.Done()
			iterate.Inc(goCount, expo-1)
		}(goCount)
	}

	wg.Wait()

}
