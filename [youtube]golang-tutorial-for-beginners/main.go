package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	bookings              = make([]UserData, 0)
)

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() { // entry point
	greetUsers()

	// for {
	// user input & check input
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTickerNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	// BOOK LOGIC
	if isValidName && isValidEmail && isValidTickerNumber {
		// call  bookTicket func
		bookTicket(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// call firstname
		firstNames := printFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}

	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidTickerNumber {
			fmt.Println("woring')")
		}

		fmt.Println("Sorry remainngticket not enough than you entered")
		// continue
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total of %v and %v still available\n", conferenceTickets, remainingTickets)
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.FirstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// var defines
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your emal: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName, lastName, email string) {
	remainingTickets -= userTickets

	// create a map for a user
	var userData = UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// MESSAGES
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName, lastName, email string) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("\n#################")
	fmt.Printf("Sending ticket %v to email address %v\n", ticket, email)
	fmt.Println("##################")

	wg.Done()
}
