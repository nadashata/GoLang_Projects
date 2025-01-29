package main

import (
	"fmt"
	"strings"
)

func greetUsers() {
	fmt.Printf("Welcome to %v Booking App\nTotal tickets available: %v, Remaining Tickets available: %v\nGet Your tickets Now\n", confernceName, confernceTickets, remainingTickets)
}

func getUserInputs() (string, string, string, uint) {
	var firstName, lastName, email string
	var userTickets uint
	fmt.Print("Enter Your First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter Your last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter Your Email address: ")
	fmt.Scan(&email)
	fmt.Print("Enter The Number of Tickets you need : ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bookTickets(firstName string, lastName string, userTickets uint, email string) {
	// only define var
	// var users map[string]string

	// create an empty map
	// var users =make(map[string]string)
	// users["firstName"] = firstName
	// users["lastName"]=lastName
	// users["email"]=email
	// users["userTickets"]=strconv.FormatUint(uint64(userTickets),10)
	//bookingUsers=append(bookingUsers,firstName +" "+lastName)

	// create struct
	var userData = userData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		noTickets: userTickets,
	}
	bookingUsers = append(bookingUsers, userData)
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("\nThank You %v %v, for booking %v Tickets for %v, you will get a confirmation email on %v \n\n%v Remaining tickets available for %v\n==============\n\n", firstName, lastName, userTickets, confernceName, email, remainingTickets, confernceName)
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookingUsers {
		///firstNames = append(firstNames, booking["firstName"]) - maps
		firstNames = append(firstNames, booking.firstName) // struct
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidNumber
}

func userMessages(isValidName bool, isValidEmail bool, isValidNumber bool) {
	fmt.Println("\nPlease Validate your inputs\n==============")
	if !isValidName {
		fmt.Println("Your First && last name must be greater than 2 charachters")
	}
	if !isValidEmail {
		fmt.Println("Your email address must contain @")
	}
	if !isValidNumber {
		fmt.Println("The Number of Tickets you can book must be > 0 and < remaining tickets")
	}
	fmt.Println("==============")
}
