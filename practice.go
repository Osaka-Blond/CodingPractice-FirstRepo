package main

import (
	"fmt"
	"strings"
)

func main() {

	// This is the introduction of the program

	var conferenceName = "GO CONFERENCE"
	const conferenceTickets = 40
	var remainingTickets = 30
	bookings := []string{}

	fmt.Printf("Hello and welcome to the %v introduction of the program\n", conferenceName)
	fmt.Printf("Tickets now are worth %v and we still have %v tickets remaining\n", conferenceTickets, remainingTickets)
	fmt.Println("I will be your instructor this evening, describing the fundamentals of this fine programming language!")

	for {
		var FirstName string
		var LastName string
		var email string
		var userTickets int

		// Information regarding tickets

		fmt.Println("Please enter first name: ")
		fmt.Scan(&FirstName)

		fmt.Println("Please enter last name: ")
		fmt.Scan(&LastName)

		fmt.Println("Please enter email address: ")
		fmt.Scan(&email)

		fmt.Println("Please enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, FirstName+" "+LastName)

		fmt.Printf("The entire slice: %v\n", bookings)
		fmt.Printf("First value: %v\n", bookings)
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you for your time %v %v! You should be receiving a confirmation email sent to %v very soon about your %v purchased tickets\n", FirstName, LastName, email, userTickets)
		fmt.Printf("There will be %v tickets remaining for the %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of the bookings are: %v\n", firstNames)

		fmt.Printf("These are all our bookings: %v\n", bookings)

	}

}
