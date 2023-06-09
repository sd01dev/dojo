package main

import "fmt"

const size = 10

func main() {
	firstCh := make(chan int, size)
	secondCh := make(chan int, size)
	outCh := make(chan int, cap(firstCh)+cap(secondCh))

	for i := 1; i <= size; i++ {
		firstCh <- i
		secondCh <- i + 10
	}
	close(firstCh)
	close(secondCh)

	fanIn(firstCh, secondCh, outCh)
	close(outCh)

	for i := range outCh {
		fmt.Println(i)
	}
}

func fanIn(ch1, ch2, outCh chan int) chan int {
	for len(ch1) != 0 || len(ch2) != 0 {
		select {
		case n := <-ch1:
			outCh <- n
		case m := <-ch2:

			outCh <- m
		}
	}

	return outCh
}
