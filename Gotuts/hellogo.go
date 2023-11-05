package main

import (
	"fmt"
)

var pl = fmt.Println

// func main() {

// 	pl("What is your name?: ")
// 	reader := bufio.NewReader(os.Stdin)

// 	name, err := reader.ReadString('\n')
// 	if err == nil {
// 		pl("Hello, ", name)
// 	} else {
// 		log.Fatal(err)
// 	}
// }

func getSumArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func varay(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return sum
}

func changVal(val *int) {
	*val += 1
	//return *val
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, val := range nums {
		sum += val
	}
	return sum / numSize
}

func arrayVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func main() {
	// val1 := "4500"
	// val2, err := strconv.Atoi(val1)

	// val1 := "4.500"
	// val2, err := strconv.ParseFloat(val1, 64)
	// pl(val2, err, reflect.TypeOf(25))

	// word1 := "A word"
	// replacer := strings.NewReplacer("A", "Another")
	// word2 := replacer.Replace(word1)

	// pl(word2)

	// pl(len(word2))
	// pl("Contains another: ", strings.Contains(word2, "Anotther"))
	// pl("o index: ", strings.Index(word2, "o"))
	// pl("Replace: ", strings.Replace(word2, "o", "0", -1))

	// word3 := "\nSome Words\n"
	// word3 = strings.TrimSpace(word3)
	// pl("Split: ", strings.Split("a-b-c-d", "-"))
	// pl("Lower: ", strings.ToLower(word2))

	// rune -> charater
	// word := "abcdefg"
	// pl("Rune count: ", utf8.RuneCountInString(word))

	// for i, runeVal := range word {
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	// }

	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())

	//seedSec := time.Now().Unix()

	// randNum := rand.Intn(50)
	// pl(randNum)

	//array := []int{1, 2, 3, 4}

	// val := 11
	// var valPtr *int = &val
	// pl("valPtr adress: ", valPtr)

	// changVal(&val)
	// pl("valPtr value func: ", *valPtr)

	// pArray := [4]int{1, 2, 3, 4}
	// arrayVals(&pArray)

	// pl(pArray)

	iSlice := []float64{12, 34, 23}
	fmt.Printf("Average: %.3f\n", getAverage(iSlice...))

}
