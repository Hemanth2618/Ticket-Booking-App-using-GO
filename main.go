package main

import (
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50

var conferenceName = "Go Conference"

/*
Syntactic sugar specific to Golang (Cannot be used when declaring a variable or initializing a constant)
conferenceName := "Go Conference
*/

var remainingTickets uint = 50

/*
	Declaring of map data structure
var bookings = make([]map[string]string, 0)
*/

// Declaring a struct data type
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// Declaring a WaitGroup to wait for completion of other goroutines
var wg = sync.WaitGroup{}

func main() {
	greetUsers()
	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)

			// Wait group is only added when the above for loop is not present. It coordinates and waits for completion of multiple goroutines
			wg.Add(1)
			// Creating a separate Go routine to achieve concurrency and parallelism
			go sendTicket(userTickets, firstName, lastName, email)
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked out. Come back next year")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or Last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// Ask user for their input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	/* create a map for the user
	var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)
	*/
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
}

// Simulating a long process (In this case, sending the booked ticket via the email client) by Sleep function
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \n to email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}

/* Switch statement example
var city string
fmt.Println("Please enter the city name")
fmt.Scan(&city)

switch city {
case "New York":
	// booking logic for New York conference
case "Singapore", "Hong Kong":
	// booking logic for Singapore and Hong Kong conference
case "London", "Berlin":
	// booking logic for London and Berlin
case "Mexico":
	// booking logic for Mexico
default:
	fmt.Println("No valid city selected")
} */
