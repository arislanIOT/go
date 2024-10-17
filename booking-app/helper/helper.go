package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, userEmail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 // var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}