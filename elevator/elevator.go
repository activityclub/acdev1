package elevator

import (
	"fmt"
	"github.com/icrowley/fake"
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

type Floor struct {
	Number   int
	WantUp   bool
	WantDown bool
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func NewFloor(number int) *Floor {
	f := Floor{}
	f.Number = number
	return &f
}

func NewPassenger(name string) *Passenger {
	p := Passenger{}
	p.Name = fake.FirstName() + " " + fake.LastName()
	p.CurrentFloor = 1
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

	for i := 1; i < 4; i++ {
		v := NewVator(i)
		go v.run()
	}

	fmt.Println("Create 30 Passengers...")

	for i := 0; i < 30; i++ {
		p := NewPassenger("")
		fmt.Println("Passenger: ", p.Name, "CurrentFloor: ", p.CurrentFloor, " DesiredFloor: ", p.DesiredFloor)
	}

	fmt.Println("Create 99 Floors...")

	floors := make([]Floor, 0)
	for i := 1; i < 100; i++ {
		f := NewFloor(i)
		floors = append(floors, *f)
	}

	fmt.Println(floors)

	wg.Wait()
}
