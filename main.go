package main

import (
	"awesomeProject/helper"
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings []string

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets, input := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			if !getAllBookings(input, bookings) {
				break
			}

			firstNames := getFirstNames()

			fmt.Printf("The first names of bookings are: %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets\n", remainingTickets)
		}
	}
}

func greetUsers() {
	fmt.Printf("---Welcome to %v booking application---\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaible\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var input string

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets, input
}

func getAllBookings(input string, bookings []string) bool {
	fmt.Println("stop or ticket")
	fmt.Scanln(&input)

	if input == "ticket" {
		for i := 0; i < len(bookings); i++ {
			fmt.Println("This is the " + bookings[i])
		}
		return true
	} else if input == "stop" {
		return false
	}
	return true
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
