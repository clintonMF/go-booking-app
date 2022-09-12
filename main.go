package main

import (
	"fmt"
	"strings"
)

var conferenceName string = "Go conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greeting()

	for {
		var firstName string
		var lastName string
		var numberOfTickets uint
		var email string

		fmt.Println("Enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&numberOfTickets)
		fmt.Println("Enter your email address")
		fmt.Scan(&email)

		isValidEmail, isValidName, isValidTicketNumber := validateUserInput(
			firstName, lastName, email, numberOfTickets)

		if isValidEmail && isValidName && isValidTicketNumber {
			//this logic is used to prevent the user from booking
			// more tickets than there are remaining

			bookTickets(firstName, lastName, email, numberOfTickets)

			firstNames := getFirstNames()
			fmt.Printf("the first names of bookings are %v \n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All our tickets have been sold out, come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("You entered a wrong email")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greeting() {
	fmt.Println("Welcome to", conferenceName, "booking site")
	fmt.Printf("We have a total of %v tickets and %v"+
		"tickets are still left", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, numOfTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := numOfTickets > 0 && numOfTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTickets(firstName string, lastName string, email string, numberOfTickets uint) {
	remainingTickets = conferenceTickets - numberOfTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a "+
		"confirmation mail at %v\n", firstName, lastName, numberOfTickets,
		email)
	fmt.Printf("%v tickets remaining for the"+
		" %v", remainingTickets, conferenceName)
}
