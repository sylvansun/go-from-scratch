package intermediate

import "fmt"

func ClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// fmt.Println("waiting for jobs")
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

/*
possible outcome of this code:
sent job 1
received job 1
received job 2
sent job 2
sent job 3
sent all jobs
received job 3
received all jobs
*/
