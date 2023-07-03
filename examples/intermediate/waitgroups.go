package intermediate

import (
	"fmt"
	"sync"
	"time"
)

func gworker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroups() {

	var wg sync.WaitGroup // Note: if a WaitGroup is explicitly passed into functions, it should be done by pointer

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			gworker(i)
		}()
	}

	wg.Wait()

}
