package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/go-routines/data"
)

func main() {
	start := time.Now()
	waitGroup := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go readBook(i, waitGroup, m)
	}

	waitGroup.Wait()
	duration := time.Since(start).Milliseconds()

	fmt.Printf("%dms\n", duration)
}

func readBook(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
	data.FinishBook(id, m)

	delay := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}
