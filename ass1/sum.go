package sum

import "fmt"

func Parallelsum(arr []int, ch chan int) {
	total := 0
	for _, value := range arr {
		total += value
	}
	ch <- total
}

func Totalsum() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ch := make(chan int, 4)
	size := 3
	for i := 0; i < len(arr)/size; i++ {
		start := i * size
		end := start + size
		go Parallelsum(arr[start:end], ch)
	}
	total := 0
	fmt.Print(ch)
	for i := 0; i < len(arr)/size; i++ {
		partialSum := <-ch
		total += partialSum
	}
	fmt.Printf("The total sum is: %d\n", total)
}
