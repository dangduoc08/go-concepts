package infextender

import (
	"fmt"
	"go-lang/inf"
)

type dog struct {
	inf.Animal
}

type animalExtender interface {
	inf.Animal
	move() string
}

func (d dog) move() string {
	return "Dog is running...!"
}

// Extend func
func Extend() {
	var aDogExtender dog = dog{inf.AnAnimal}
	var anAnimalExtend animalExtender = aDogExtender

	var dogMove string = anAnimalExtend.move()
	fmt.Println(dogMove)

	var dogYield string = anAnimalExtend.Yield()
	fmt.Println(dogYield, "In extend package")
}
