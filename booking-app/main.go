// Conference Booking Application //
package main

import (
	"fmt"
	// "strings"
	"booking-app/helper"
	// "strconv"
	"time"
)

var remainingTickets uint = 50
var bookings = make([]UserData, 0) // we need to initilize the size for the map of slice which will grwo auto

// strcut gives us power to define what user inputs will be ...
// strcut can be compared to classes in java etc..
type UserData struct{
	firstName string
	lastName string
	email string
	userTickets uint
	isOptedForInNews bool
}

func main(){
	// fmt.Println("Welcome to Conference Booking App")
	// fmt.Println("Get your tickets here to attend")
	

	for {											//indefinete loop

		firstName, lastName, userTickets, userEmail := getUserInput()
		
		

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userEmail, userTickets, remainingTickets) // validate user input function

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(firstName, lastName, userTickets, userEmail)
			go sendTicket(firstName, lastName, userTickets, userEmail)
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
		// firstnames = append(firstnames, bookings["firstName"])	for map
		firstnames = append(firstnames, bookings.firstName)	
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


	// var userData = make(map[string]string)

	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["Email"] = userEmail
	// userData["numberofTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: userEmail,
		userTickets: userTickets,
	}

	// Now we need to add these userData into booking. 
	// To do that we need to make booking into a slice of map.

	bookings = append(bookings, userData)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for choosing Confa; your %v tickets shared to your email %v\n", firstName, lastName, userTickets, userEmail)
	fmt.Printf("Remaining Tickets: %v\n", remainingTickets)
	// return bookings, remainingTickets, userTickets
	// return remainingTickets
}

func sendTicket(firstName string, lastName string, userTickets uint, userEmail string){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("Sending ticket:\n %v \nto the email address %v\n", ticket, userEmail)

}