package main

import (
	"fmt"
)

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Insurance interface {
	CalculateInsurance() float64
}

type Printable interface {
	Details() string
}

type Car struct {
	Vehicle
	NumberOfDoors int
}

func (car Car) Details() string {
	return fmt.Sprintf("Car: %s, Model: %s, Year: %d, Number of Doors: %d ", car.Make, car.Model, car.Year, car.NumberOfDoors)
}

func (car Car) CalculateInsurance() float64 {
	age := float64(2025-car.Year) * 10.0
	return age
}

type Truck struct {
	Vehicle
	PayloadCapacity float64
}

func (truck Truck) Details() string {
	return fmt.Sprintf("Truck : %s, model : %s, Year: %d, Capacity: %f", truck.Make, truck.Model, truck.Year, truck.PayloadCapacity)
}

func (truck Truck) CalculateInsurance() float64 {
	age := float64(2025-truck.Year) * 10.0
	payloadFactor := truck.PayloadCapacity * 100.0
	return age + payloadFactor
}

func print(printables []Printable) {
	for _, p := range printables {
		fmt.Println(p.Details())
	}
}

func main() {

	car1 := Car{Vehicle: Vehicle{Make: "Toyota", Model: "Corolla", Year: 2020}, NumberOfDoors: 4}
	car2 := Car{Vehicle: Vehicle{Make: "RollsRoyce", Model: "cullinan", Year: 2022}, NumberOfDoors: 4}
	car3 := Car{Vehicle: Vehicle{Make: "Mercedes", Model: "maybach", Year: 2024}, NumberOfDoors: 4}
	truck1 := Truck{Vehicle: Vehicle{Make: "Ford", Model: "model1", Year: 2018}, PayloadCapacity: 3.5}
	truck2 := Truck{Vehicle: Vehicle{Make: "Mercedes", Model: "model2", Year: 2020}, PayloadCapacity: 5}
	truck3 := Truck{Vehicle: Vehicle{Make: "Ford", Model: "model3", Year: 2020}, PayloadCapacity: 4.5}

	printables := []Printable{car1, car2, car3, truck1, truck2, truck3}

	print(printables)

}
