package main

import (
	"sync"
	"time"
	"fmt"
)	

var conferenceName = "Go Conference"
const conferenceTickets = 50;
var remaininigTickets uint= 50
var bookings = make([]userData, 0)

type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
	
}

var wg = sync.WaitGroup{}

func main(){
	
	greetUser()



		firstName,lastName,userTickets,email := getUserInput()

		isValidName,isValidEmail,isValidTicketNumber := validateUserInput(firstName,lastName,email,userTickets,remaininigTickets)

		

	if isValidName && isValidEmail && isValidTicketNumber{
		bookTickets(userTickets,firstName,lastName,email)

		wg.Add(1)
		go sendTicket(userTickets,firstName,lastName,email)
		
	    // call function Print firstnames
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)


	    if remaininigTickets == 0{
		fmt.Println("Our conference is booked out. Come back next year")
		// break
	    }

	} else {
		if !isValidName{
			fmt.Println("First name or Last name you entered is too short.")
		}
		if !isValidEmail{
			fmt.Println("Email address you entered is not valid.")
		}
		if !isValidTicketNumber{
			fmt.Println("Number of tickets you entered is invalid.")
		}
		
		
	}
	wg.Wait()
}
func greetUser(){
	fmt.Printf("Welcome to booking %v application \n", conferenceName)
	fmt.Printf("We hare total of %v tickets and %v are still available.\n", conferenceTickets, remaininigTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames()[]string{
	firstNames := []string{}
	    for _, booking := range bookings {
		   firstNames = append(firstNames,booking.firstName)
	    }
		return firstNames
	    
}

func getUserInput()(string,string,uint,string){
	    var firstName string
	    var lastName string
	    var userTickets uint
	    var email string
	    fmt.Println("Enter your first name: ")
	    fmt.Scan(&firstName)

	    fmt.Println("Enter your last name: ")
	    fmt.Scan(&lastName)
	
	    fmt.Println("Enter your  emain address: ")
	    fmt.Scan(&email)

	    fmt.Println("Enter number of tickets: ")
	    fmt.Scan(&userTickets)
		return firstName,lastName, userTickets,email
}

func bookTickets(userTickets uint,firstName string,lastName string,email string){
	    remaininigTickets = remaininigTickets - userTickets

		var userData = userData{
			firstName: firstName,
			lastName: lastName,
			email: email,
			numberOfTickets: userTickets,
		}

	    bookings = append(bookings, userData)

		fmt.Printf("List of bookings is %v\n", bookings)

	    fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTickets, email)

	    fmt.Printf("%v tickets remaining for %v\n", remaininigTickets, conferenceName) 
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50*time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket,email)
	fmt.Println("*****************")

wg.Done()

}