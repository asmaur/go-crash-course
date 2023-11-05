package main

import (
	"fmt"
)

var pl = fmt.Println

type customer struct {
	name    string
	address string
	bal     float64
}

type Rectangle struct {
	length, height float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

func getCustomerInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

func main() {
	tS := customer{"Maur", "5 main st", 12.34}

	getCustomerInfo(tS)
	newCustAdd(&tS, "tets and")
	getCustomerInfo(tS)
	pl(tS.address)

	rect1 := Rectangle{10.5, 16.78}
	pl("Rect area: ", rect1.Area())

}
