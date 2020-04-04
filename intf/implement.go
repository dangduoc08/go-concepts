package intf

import "fmt"

type dog struct{}

// Animal interface
type Animal interface {
	Yield() string
}

func (d dog) Yield() string {
	return "Gau gau...!"
}

var aDog dog = dog{}

// AnAnimal var
var AnAnimal Animal = aDog

// Implement func
func Implement() {
	var dogYield string = AnAnimal.Yield()
	fmt.Println(dogYield)
}
