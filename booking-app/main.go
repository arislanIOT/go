// Conference Booking Application //
package main

import (
	"fmt"
	"strings"
)

var remainingTickets uint = 50
var bookings []string

func main(){
	// fmt.Println("Welcome to Conference Booking App")
	// fmt.Println("Get your tickets here to attend")
	

	for {											//indefinete loop

		firstName, lastName, userTickets, userEmail := getUserInput()
		
		

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets) // validate user input function

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, userTickets, userEmail)
			firstnames := getFirstName(bookings)// call the fucntion to print the bookings name
			fmt.Printf("These are all your bookings: %v\n", firstnames)

			if remainingTickets == 0 {
				fmt.Println("Thank you for choosing confa, our ticket sold out")
				break
			}

		} else {

			if !isValidName { 
				fmt.Println("The first or lastname is two short; min 2 required")
			}

			if !isValidEmail {
				fmt.Println("Missing @ in the email")
			}
			if !isValidTicketNumber {
				fmt.Println("Enter valid ticket number")
			}
	    }

	}

}

func getFirstName(bookings []string) []string {
	firstnames := []string{}
	for _, bookings := range bookings {
		var names = strings.Fields(bookings)
		firstnames = append(firstnames, names[0])	
	}
	//fmt.Printf("These are all your bookings: %v\n", firstnames)
	return firstnames
}
func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 // var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}
func getUserInput() (string, string, uint, string) {
	var firstName string
	var lastName string
	var userTickets uint
	var userEmail string 
	// bookings := []string{}
	fmt.Println("Welcome to Confa, Please state your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Welcome to Confa, Please state your last name: ")
	fmt.Scan(&lastName)


	fmt.Println("Kindly enter the number of ticket required: ")
	fmt.Scan(&userTickets)
	
	fmt.Println("Kindly provide the email address: ")
	fmt.Scan(&userEmail)

	return firstName, lastName, userTickets, userEmail
}
func bookTicket(firstName string, lastName string, userTickets uint, userEmail string) {
	bookings = append(bookings, firstName + " " + lastName)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for choosing Confa; your %v tickets shared to your email %v\n", firstName, lastName, userTickets, userEmail)
	// return bookings, remainingTickets, userTickets
	// return remainingTickets
}