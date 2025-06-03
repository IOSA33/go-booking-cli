package main

import (
	"awesomeProject/helper"
	"fmt"
	"strconv"
)

// variables
const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"

// map for user information, like firstname, lastname etc
var bookings = make([]map[string]string, 0)

func main() {
	// Printing hello to user
	greetUsers()

	// while loop for booking tickets
	for {
		// getting user information from inputs
		firstName, lastName, email, userTickets, input := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		// check if any of values is available
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

// saying hello to user
func greetUsers() {
	fmt.Printf("---Welcome to %v booking application---\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still avaible\n", conferenceTickets, remainingTickets)
}

// getting firstNames of bookings map (List) by key "firstName"
func getFirstNames() []string {
	firstNames := []string{}
	// _ is index and booking is value of the index
	for _, booking := range bookings {
		var names = booking["firstName"]
		firstNames = append(firstNames, names)
	}
	return firstNames
}

// asking user for information via inputs
func getUserInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var input string

	fmt.Println("Enter Your First Name: ")
	_, err := fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	_, err = fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	_, err = fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	_, err = fmt.Scanln(&userTickets)

	// handle error if example input is null
	if err != nil {
		fmt.Println("Input error:", err)
	}
	return firstName, lastName, email, userTickets, input
}

// printing all bookings from all users
func getAllBookings(input string, bookings []map[string]string) bool {
	fmt.Println("stop or ticket")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Input error in stop or ticket:", err)
		return true
	}

	if input == "ticket" {
		for i := 0; i < len(bookings); i++ {
			fmt.Println("This is the ", bookings[i])
		}
		return true
	} else if input == "stop" {
		return false
	}
	return true
}

// booking tickets for one user with information
func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create map for a user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)

	fmt.Printf("List of bookings: %v \n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
