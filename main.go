package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Go conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]userData, 0)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greeting()

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

	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(
		firstName, lastName, email, numberOfTickets)

	if isValidEmail && isValidName && isValidTicketNumber {
		//this logic is used to prevent the user from booking
		// more tickets than there are remaining

		bookTickets(firstName, lastName, email, numberOfTickets)
		wg.Add(1) // this function sets the number of goroutines to wait for
		go sendTicket(numberOfTickets, email, firstName, lastName)

		firstNames := getFirstNames()
		fmt.Printf("the first names of bookings are %v \n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All our tickets have been sold out, come back next year")
			// break
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
	wg.Wait() // this function blocks the main thread until the WaitGroup counter is 0
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
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func bookTickets(firstName string, lastName string, email string, numberOfTickets uint) {
	remainingTickets = conferenceTickets - numberOfTickets
	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("list of bookings are => %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a "+
		"confirmation mail at %v\n", firstName, lastName, numberOfTickets,
		email)
	fmt.Printf("%v tickets remaining for the %v\n", remainingTickets, conferenceName)
}

func sendTicket(numberOfTickets uint, email string, firstName string, lastName string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)

	fmt.Println("######################")
	fmt.Printf("Sending tickets:\n%v \nto email address :%v\n", ticket, email)
	fmt.Println("######################")
	wg.Done() //this function decrements the wait group counter by 1
}
