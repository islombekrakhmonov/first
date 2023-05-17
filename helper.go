package main

import "strings"

func 	validateUserInput(firstName string, lastName string, email string, userTickets uint, remaininigTickets uint)(bool,bool,bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets >0 && userTickets <= remaininigTickets
	return isValidName,isValidEmail,isValidTicketNumber
}