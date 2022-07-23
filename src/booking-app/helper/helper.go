package helper

import "strings"

// function to validate user input
// returns 3 booleans
// capitalize first letter to export function
func ValidateUserInput(firstName string, lastName string, email string, ticketAmount uint, remainingTickets uint) (bool, bool, bool) {
	// bool var for name validity
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// bool var for email validity
	isEmailValid := strings.Contains(email, "@")
	// bool var for tickets amount validity
	isTicketsAmountValid := ticketAmount > 0 && ticketAmount <= remainingTickets

	return isValidName, isEmailValid, isTicketsAmountValid
}