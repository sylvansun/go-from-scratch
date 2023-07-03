package intermediate

import "fmt"

func RangeOverChannels() {

	queue := make(chan string, 3)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
