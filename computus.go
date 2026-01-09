package computus

import "time"

// Easter returns the date of Easter Sunday for the given year
// using the Anonymous Gregorian Algorithm.
// ref: https://en.wikipedia.org/wiki/Date_of_Easter#Anonymous_Gregorian_algorithm
func Easter(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451

	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

// SundayLetters returns the Sunday letter(s) for a given year.
// If the year is not a leap year, only one letter is returned.
// If the year is a leap year, two letters are returned:
// - first applies Jan 1 – Feb 24
// - second applies Feb 25 – Dec 31
func SundayLetters(year int) (string, string) {
	const letters = "ABCDEFG"

	// January 1 of the year
	jan1 := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Weekday: Sunday=0 ... Saturday=6
	weekday := int(jan1.Weekday())

	// Sunday letter for Jan 1
	letter := letters[(7-weekday)%7 : (7-weekday)%7+1]

	if isLeapYear(year) {
		// In leap years, after Feb 24, the letter shifts backward by one
		// e.g., if first letter = A, second letter = G
		idx := (7 - weekday - 1 + 7) % 7
		secondLetter := letters[idx : idx+1]
		return letter, secondLetter
	}

	return letter, ""
}

// isLeapYear returns true if the year is a Gregorian leap year
func isLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	} else if year%100 != 0 {
		return true
	} else if year%400 != 0 {
		return false
	}
	return true
}
