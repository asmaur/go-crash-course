package main

import "fmt"

func main() {
	ids := []int{34, 56, 12, 36, 78, 32}

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID: %d\n", i, id)
	// }

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0

	for _, id := range ids {
		sum += id
	}
	fmt.Printf("Sum: %d\n", sum)

	// range maps
	emails := map[string]string{"bob": "bob@gmail.com", "kurbi": "kurb@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s \n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name : %s \n", k)
	}
}
