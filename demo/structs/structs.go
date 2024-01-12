package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	cassey := Passenger{"Cassey", 23, false}
	fmt.Println(cassey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 25}
		ella = Passenger{Name: "Ella", TicketNumber: 28}
	)
	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumber = 21
	fmt.Println(heidi)

	cassey.Boarded = true
	bill.Boarded = true

	if bill.Boarded {
		fmt.Println("Bill has boarded the bus")
	}
	if cassey.Boarded {
		fmt.Println(cassey.Name, "has boarded the bus")
	}

	heidi.Boarded = true

	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is on the front seat.")
}
