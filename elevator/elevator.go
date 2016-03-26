package elevator

import (
	"fmt"
	"github.com/icrowley/fake"
	"math/rand"
	"sync"
	"time"
)

var MAX_FLOORS int = 99

type Vator struct {
	Number       int
	CurrentFloor int
	NextFloor    int
	Buttons      []bool
	Passengers   []*Passenger
}

type Passenger struct {
	Name         string
	CurrentFloor int
	DesiredFloor int
}

type Floor struct {
	Number     int
	WantUp     bool
	WantDown   bool
	m          sync.Mutex
	Passengers []*Passenger
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func NewFloor(number int) *Floor {
	f := Floor{}
	f.Number = number
	f.Passengers = make([]*Passenger, 0)
	return &f
}

func NewPassenger(name string) *Passenger {
	p := Passenger{}
	p.Name = fake.FirstName() + " " + fake.LastName()
	p.CurrentFloor = 1
	p.DesiredFloor = random(1, MAX_FLOORS)

	return &p
}

func NewVator(number int) *Vator {
	v := Vator{}
	v.Number = number
	v.CurrentFloor = 1
	v.Buttons = make([]bool, MAX_FLOORS)
	v.Passengers = make([]*Passenger, 0)

	return &v
}

func (self *Vator) run(floors []*Floor) {
	for {
		fmt.Println("Vator ", self.Number, " CurrentFloor: ", self.CurrentFloor)
		time.Sleep(time.Second * 5)
	}
}

func (self *Floor) upButton() {
	if self.WantUp {
		return
	}
	fmt.Printf("  UP button on Floor %d pushed,\n", self.Number)
	self.WantUp = true
}

func passengerCreator(floor1 *Floor) {
	for {
		p := NewPassenger("")
		floor1.m.Lock()
		floor1.Passengers = append(floor1.Passengers, p)
		floor1.m.Unlock()
		fmt.Printf("New Passenger (%s) just walked in on Floor 1.\n", p.Name)
		fmt.Printf("  They want to be on Floor %d \n", p.DesiredFloor)
		floor1.upButton()
		time.Sleep(time.Second * time.Duration(random(0, 10)))
	}
}

func Run() {
	rand.Seed(time.Now().Unix())
	var wg sync.WaitGroup

	wg.Add(1)

	fmt.Println("Create Floors...")

	floors := make([]*Floor, 0)
	for i := 1; i < MAX_FLOORS+1; i++ {
		f := NewFloor(i)
		floors = append(floors, f)
	}

	fmt.Println("Create 3 Elevators...")

	for i := 1; i < 4; i++ {
		v := NewVator(i)
		go v.run(floors)
	}
	go passengerCreator(floors[0])

	wg.Wait()
}
