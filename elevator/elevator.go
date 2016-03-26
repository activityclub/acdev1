package elevator

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Vator struct {
	Number       int
	CurrentFloor int
	NextFloor    int
	Passengers   int
}

type Passenger struct {
	Name         string
	CurrentFloor int
	DesiredFloor int
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func NewPassenger(name string) *Passenger {
	p := Passenger{}
	p.Name = "John Smith"
	p.CurrentFloor = random(1, 99)
	p.DesiredFloor = random(1, 99)

	return &p
}

func NewVator(number int) *Vator {
	v := Vator{}
	v.Number = number
	v.CurrentFloor = 1
	v.Passengers = 0

	return &v
}

func (self *Vator) run() {
	for {
		fmt.Println("Vator ", self.Number, " CurrentFloor: ", self.CurrentFloor)
		time.Sleep(time.Second * 5)
	}
}

func Run() {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup

	wg.Add(1)
	fmt.Println("Create 3 Elevators...")

	i := 1
	for {
		if i > 3 {
			break
		}

		v := NewVator(i)
		go v.run()

		i++

	}

	fmt.Println("Create 30 Passengers...")

	i = 0
	for {
		if i > 30 {
			break
		}

		p := NewPassenger("")
		fmt.Println("Passenger: ", p.Name, "CurrentFloor: ", p.CurrentFloor, " DesiredFloor: ", p.DesiredFloor)

		i++

	}

	wg.Wait()
}
