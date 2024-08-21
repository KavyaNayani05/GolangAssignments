/*package main

import (
	"fmt"
	subtract "golang/assignment"
)

var name string = "golang practice"

func main() {
	a, b := 50, 20
	fmt.Println(name)
	if a > b {
		subtract.Sub(a, b)
	} else {
		subtract.Sub(b, a)
	}
	var i = 0
	for i < 5 {
		fmt.Println("i value is", i)
		i++
	}
	//arrays
	arr := [4]int{2, 3, 4, 5}
	arr[2] = 11
	fmt.Println("Array is", arr)
	fmt.Println("Length of an array is", len(arr))
	//slices
	var names = []string{"apple", "mango", "peach"}
	names[1] = "strawberry"
	names = append(names, "kiwi")
	fmt.Println("After appending", names)
	var new_arr = arr[1:3]
	fmt.Println("New array is", new_arr)
	//maps
	scores := map[string]float64{
		"Emma":   89.5,
		"Ashley": 90.2,
		"Robert": 78.4,
		"Peter":  92.3,
	}
	fmt.Println(scores)
	scores["Robert"] = 80.0
	fmt.Println(scores)
	delete(scores, "Emma")
	for k, v := range scores {
		fmt.Println(k, "-", v)
	}
}
*/

package main

import "fmt"

func primes(n int, ch chan int) {
	for i := 2; i <= n; i++ {
		isPrime := true
		for p := 2; p*p <= i; p++ {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			ch <- i
		}
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go primes(50, ch)
	//x := len(ch)
	//fmt.Println(x)
	for value := range ch {
		fmt.Println(value)
	}
}
