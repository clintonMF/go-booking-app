package main

import "strings"

/*
Todo: turn this file to a package
to ensure you got the concept of creating packages in golang
*/
func validateUserInput(firstName string, lastName string, email string,
	numOfTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := numOfTickets > 0 && numOfTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
