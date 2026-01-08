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

// AshWednesday returns the date of Ash Wednesday for the given year.
// Ash Wednesday is 46 days before Easter Sunday.
func AshWednesday(year int) time.Time {
	easter := Easter(year)
	return easter.AddDate(0, 0, -46)
}

// PalmSunday returns the date of Palm Sunday for the given year.
// Palm Sunday is always 7 days before Easter Sunday.
func PalmSunday(year int) time.Time {
	easter := Easter(year)
	return easter.AddDate(0, 0, -7)
}

// GoodFriday returns the date of Good Friday for the given year.
// Good Friday is always 2 days before Easter Sunday.
func GoodFriday(year int) time.Time {
	easter := Easter(year)
	return easter.AddDate(0, 0, -2)
}
