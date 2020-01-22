package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log(counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {

				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Log(counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	wg.Add(1)
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {

				mut.Unlock()
			}()
			mut.Lock()
			wg.Wait()
			counter++
		}()
	}
	wg.Done()
	t.Log(counter)
}
