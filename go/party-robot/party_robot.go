package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
	s := Welcome(name)
	s += fmt.Sprintf("\nYou have been assigned to table %X. ", table)
	s += fmt.Sprintf("Your table is %s, ", direction)
	s += fmt.Sprintf("exactly %.1f meters from here.", distance)
	s += fmt.Sprintf("\nYou will be sitting next to %s.", neighbour)
	return s
}
