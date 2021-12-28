package helper

import "net/mail"

func validMailAddress(address string) bool {
	_, err := mail.ParseAddress(address)

	if err != nil {
		return false
	}

	return true
}

func ValidUserInput(firstName, lastName, email string, userTickets, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := validMailAddress(email)
	isValidAmountOfTickets := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidAmountOfTickets
}
