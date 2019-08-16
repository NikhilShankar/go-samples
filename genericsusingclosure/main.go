package main

import (
"fmt"
)

type Car struct {
	Name string
	Id   int
}

type Plane struct {
	Name string
	Id   int
}

//This is the crux of the idea.
// Creating generics using closures.
// Refer this youtube video by Jon Bodner for more info
// https://www.youtube.com/watch?v=5IKcPMJXkKs
func filter(len int,searchName string, getNameFunction func(i int) string, foundAt func(i int)) {
	for i := 0; i < len; i++ {
		if getNameFunction(i) == searchName {
			foundAt(i)
		}
	}
}

func main() {
	fmt.Println("Hello, playground")
	car1 := Car {
		Name : "swift",
		Id : 1,
	}

	car2 := Car {
		Name : "benz",
		Id : 2,
	}

	cars := []Car{car1, car2}


	plane1 := Plane {
		Name : "boeing",
		Id : 1,
	}

	plane2 := Plane {
		Name : "boomboom",
		Id : 2,
	}

	planes := []Plane{plane1, plane2}


	var c *Car

	filter(len(cars),"swift", func(i int) string {
		return cars[i].Name
	}, func(i int) {
		c = &cars[i]
	})

	if c != nil {
		fmt.Println("The car is found" , c.Name)
	}


	var p *Plane

	filter(len(planes),"boomboo", func(i int) string {
		return planes[i].Name
	}, func(i int) {
		p = &planes[i]
	})

	if p != nil {
		fmt.Println("The car is found" , p.Name)
	}
}

