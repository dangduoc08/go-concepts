package intfextender

import (
	"fmt"
	"golang-concepts/intf"
)

type dog struct {
	intf.Animal
}

type animalExtender interface {
	intf.Animal
	move() string
}

func (d dog) move() string {
	return "Dog is running...!"
}

// Extend func
func Extend() {
	var aDogExtender dog = dog{intf.AnAnimal}
	var anAnimalExtend animalExtender = aDogExtender

	var dogMove string = anAnimalExtend.move()
	fmt.Println(dogMove)

	var dogYield string = anAnimalExtend.Yield()
	fmt.Println(dogYield, "In extend package")
}
