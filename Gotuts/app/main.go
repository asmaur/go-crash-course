package main

import (
	"fmt"
	"log"

	stuff "github.com/asmaur/go_crash_course/Gotuts/app/mypackage"
)

func main() {
	// fmt.Println("Heelo: ", stuff.Name)
	// fmt.Println(stuff.IntArrayToStrArr([]int{1, 2, 3, 4, 5}))

	date := stuff.Date{}
	err := date.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}

	err = date.SetYear(1890)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("1st Day: %d/%d/%d", date.Day(), date.Month(), date.Year())
}
