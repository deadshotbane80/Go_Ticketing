package helper

import "fmt"

func GreetUsers(conferenceName string, conferenceTickets int, remainingTickets uint) {

	fmt.Printf("Welcome to %v booking application\n", conferenceName) //%v is for formatting
	fmt.Printf("We have total of %v tickets and the remaining tickets are %v\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
