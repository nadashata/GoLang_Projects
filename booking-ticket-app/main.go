package main

import (
	"fmt"
)

// package level Var we cannot use := syntax

var confernceName = "Go Conference"

const confernceTickets = 50

var remainingTickets uint = 50

//var bookingUsers []string
//var bookingUsers= make([]map[string]string, 0) // create empty list of maps, it is still slice but we need to define min size which is 0

type userData struct {
	firstName string
	lastName  string
	email     string
	noTickets uint
}

var bookingUsers = make([]userData, 0)

func main() {
	greetUsers()
	for remainingTickets > 0 {
		firstName, lastName, email, userTickets := getUserInputs()
		isValidName, isValidEmail, isValidNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidNumber {
			bookTickets(firstName, lastName, userTickets, email)
			firstNames := getFirstNames()
			fmt.Printf("FirstNames of People who managed to book:%v", firstNames)
		} else {
			userMessages(isValidName, isValidEmail, isValidNumber)
		}
	}
}

//////////////////////////////////////////////////////////////////////////////////
