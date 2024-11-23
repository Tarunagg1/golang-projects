package main

import (
	"fmt"
)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickkets uint
}

func main() {

	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var conferenceName = "Go Conference"
	var bookings []string
	valudateuserInput()

	fmt.Println("Welcome to ", conferenceName, " Booking System")
	fmt.Println("We have total ", conferenceTickets, " tickets available for this event is: ", remainingTickets)
	fmt.Println("Get Your tickets here to attend")
	for {

		var firstName string
		fmt.Println("Enter your firstName")
		fmt.Scan(&firstName)

		var lastName string
		fmt.Println("Enter your lastName")
		fmt.Scan(&lastName)

		var email string
		fmt.Println("Enter your email")
		fmt.Scan(&email)

		var userTickkets uint
		fmt.Println("Enter your userTickkets")
		fmt.Scan(&userTickkets)

		remainingTickets -= userTickkets

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Println("Thankyou %v %v for booking. you will receive confimation email at %v againts tickets count %v\n", firstName, lastName, email, userTickkets)

		if remainingTickets == 0 {
			fmt.Println("event all Tickets are sold")
			break
		}
	}
}
