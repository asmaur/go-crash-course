package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	regex := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?", regex)
	pl(match)
}
