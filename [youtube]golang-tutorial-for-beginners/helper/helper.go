package helper

import "strings"

func ValidateUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickerNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTickerNumber
}
