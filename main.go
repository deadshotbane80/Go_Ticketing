package main //import package in go

import (
	"booking-app/helper"
	"fmt" //import fmt package to use print,formatiing,input,etc(Refer the go documentation)
	"time"

	"strings"
)

var conferenceName = "Go Conference" //variable name
const conferenceTickets = 50         //constant variable
var remainingTickets uint = 50       //unint is to assign only positive integers
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName  string
	email     string
	userTix   uint
}

func main() { //main function in go

	helper.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for { //infinite for loop
		var firstName string //ask user for the name
		var lastName string
		var email string
		var userTix uint
		//Note in go we need a pointer to take input
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		validName := len(firstName) <= 2 || len(lastName) <= 2
		if validName {
			fmt.Println("Enter your first name and last name greater than 2 characters")
			continue
		}
		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)
		validEmail := strings.Contains(email, "@")
		if !validEmail {
			fmt.Println("Enter valid email address")
			continue
		}
		fmt.Println("Enter no of tickets: ")
		fmt.Scan(&userTix)
		if userTix <= remainingTickets {

			book(remainingTickets, userTix, bookings, firstName, lastName, email)
			go sendTix(userTix, firstName, lastName, email)

		} else {
			fmt.Printf("You only have %v tickets remaining please enter valid no of tickets\n", remainingTickets)

		}

		if remainingTickets == 0 {
			fmt.Println("All our tickets are booked out sorry come next year")
			break
		}

	}

}

func first(bookings []userData) []string {
	var firstNames []string
	for _, booking := range bookings { //looping through a list is also known as for each loop

		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func book(remainingTickets uint, userTix uint, bookings []userData, firstName string, lastName string, email string) {

	var userAdd = userData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		userTix:   userTix,
	}

	remainingTickets = remainingTickets - userTix
	/*fmt.Println(remainingTickets)  This will print the remaining tickets
	fmt.Println(&remainingTickets) This will print out the memory location of remaining tickets
	*/
	bookings = append(bookings, userAdd) //In slices to update we append the slice
	/*fmt.Printf("Whole Slice: %v\n", bookings)
	fmt.Printf("First booking: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))
	The above code is just to better understand slicing
	*/
	fmt.Printf("List of bookings:  %v\n", bookings)
	fmt.Printf("First name of the bookings are %v\n", first(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets.You will get a confirmation email on %v\n", firstName, lastName, userTix, email)
	fmt.Printf("%v tickets are remaining\n", remainingTickets)
}

func sendTix(userTix uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) //to simulate time delay
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTix, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("Sending: \n %v to email address %v\n", ticket, email)
	fmt.Println("######################")
}

//Published on github
