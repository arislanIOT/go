// Conference Booking Application //
package main

import (
	"fmt"
	// "strings"
	"booking-app/helper"
	"strconv"
)

var remainingTickets uint = 50
var bookings = make([]map[string]string, 0) // we need to initilize the size for the map of slice which will grwo auto


func main(){
	// fmt.Println("Welcome to Conference Booking App")
	// fmt.Println("Get your tickets here to attend")
	

	for {											//indefinete loop

		firstName, lastName, userTickets, userEmail := getUserInput()
		
		

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets) // validate user input function

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, userTickets, userEmail)
			firstnames := getFirstName()// call the fucntion to print the bookings name
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

func getFirstName() []string {
	firstnames := []string{}
	for _, bookings := range bookings {
		firstnames = append(firstnames, bookings["firstName"])	
	}
	//fmt.Printf("These are all your bookings: %v\n", firstnames)
	return firstnames
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


	var userData = make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["Email"] = userEmail
	userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	// Now we need to add these userData into booking. 
	// To do that we need to make booking into a slice of map.

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for choosing Confa; your %v tickets shared to your email %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("User Data %v\n", bookings)
	// return bookings, remainingTickets, userTickets
	// return remainingTickets
}