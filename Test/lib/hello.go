package lib

import (
	"fmt"
	"time"
)

type Animal struct {
	name, breed string
}

type Lion struct {
	Animal
	name string
}

func main124() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())
	ref := 1234

	one(&ref)

	var animal = Animal{"Tommy", "Dog"}
	// var animal=Animal{name:"Tommy",breed:"Dog"};
	var defaultanimal = new(Animal)
	// fmt.Println("default animal object "+*animal)
	fmt.Println("default animal name  is " + defaultanimal.name)
	fmt.Println("Animal name  is " + animal.name)
	fmt.Println("Animal name  is " + animal.breed)
	animal.animalDescription()
	animal.animalName()

	var lion = Lion{animal, "Lion"}
	lion.Animal.roars()
}

func (animal *Animal) animalDescription() {
	fmt.Println("Animal Description for :" + animal.name)
}

func (animal *Animal) animalName() {
	fmt.Println("Animal Name is :" + animal.name)
}

func one(ref *int) {
	// *ref = 123
	var scanValue int32
	fmt.Println("Ref value is :", *ref)
	fmt.Scan(&scanValue)
	fmt.Println("Scanned value is :", scanValue)
	fmt.Println("Adress using & value is :", &scanValue)
}

func (animal *Animal) roars() {
	fmt.Println("Animal roars" + animal.name)
}
