package main

import "fmt"

func main() {

	//emails := make(map[string]string)

	// emails["bob"] = "bob@gmail.com"
	// emails["kurbi"] = "kurb@gmail.com"

	emails := map[string]string{"bob": "bob@gmail.com", "kurbi": "kurb@gmail.com"}

	fmt.Println(emails["bob"])

	//delete
	delete(emails, "bob")
	fmt.Println(emails)
}
