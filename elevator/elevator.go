package elevator

import (
	"fmt"
	"sync"
	"time"
)

type Vator struct {
	Number       int
	CurrentFloor int
	NextFloor    int
	Passengers   int
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
		fmt.Println("Vator ", self.Number)
		time.Sleep(time.Second * 5)
	}
}

func Run() {
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

	wg.Wait()
}
