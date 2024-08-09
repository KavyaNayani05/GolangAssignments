package subtract

import "fmt"

func Sub(a, b int) {
	fmt.Printf("Difference of %v and %v is %v\n", a, b, a-b)
	switch a {
	case 10:
		fmt.Println("Value of a is", 10)
	case 20:
		fmt.Println("Value of a is", 20)
	case 30:
		fmt.Println("Value of a is", 30)
	default:
		fmt.Println("Value of a is greater than 30")
	}
}
