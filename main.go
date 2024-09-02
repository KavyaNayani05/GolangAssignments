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

//goroutines and channels
/*package main

import (
	"fmt"
	sum "golang/ass1"
	"time"
)

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
	go sum.Totalsum()
	time.Sleep(2 * time.Second)
}

*/

/*package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "sending")
			ch <- i
			fmt.Println(time.Now(), i, "sent")
		}

		fmt.Println(time.Now(), "all completed")
	}()
	time.Sleep(2 * time.Second)
	fmt.Println(time.Now(), "waiting for messages")
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "received", <-ch)
	fmt.Println(time.Now(), "exiting")
}
*/

//reading input from the user

/*package main

import (
	"fmt"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Give input name ")
	//name, _ := reader.ReadString('\n')
	//name = strings.TrimSpace(name)
	//fmt.Printf("the input name given by the user is %v", name)
	var name string
	fmt.Println("Enter name ")
	fmt.Scanf("%s", &name)
	//(or) fmt.Scan(&name) or fmt.Scanln(&name)
	x := fmt.Sprintln(name)
	fmt.Println("name given by the user is ", x)
}
*/

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang/api"
	"log"
	"net/http"
)

func main() {
	dsn := "root:Kavya05@tcp(localhost:3306)/Library"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening database", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("Error connecting to database", err)
	}
	fmt.Println("Database connection was successful")
	api.RegisterRouter(db)

	log.Println("serever strt on port 8080:")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
