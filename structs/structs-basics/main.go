package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name, lastName string
	age            int
}

func main() {
	var people = []Person{
		{"Pablo", "Picasso", 91},
		{"Freud", "Sigmund", 83},
		{"Urho", "Kekkonen", 96},
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go totalAge(people, &wg)
	go listPeople(people, &wg)
	wg.Wait()
}

func listPeople(person []Person, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, person := range person {
		time.Sleep(time.Nanosecond * 400)
		fmt.Printf("My name is %s %s and my age is %d, Bye!\n", person.name, person.lastName, person.age)
	}
}

func totalAge(people []Person, wg *sync.WaitGroup) {
	defer wg.Done()
	var totalAge int
	for _, person := range people {
		totalAge = totalAge + person.age
	}
	fmt.Printf("Total age of all the people is %d \n", totalAge)
}
