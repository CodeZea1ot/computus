package computus

import "time"

// RelativeToEaster represents a moveable feast or fast whose date is relative to Easter
type RelativeToEaster struct {
	Name   string // Name of the feast/fast
	Offset int    // Days relative to Easter (negative = before, positive = after)
}

// RelativeToEasterDays represents a collection of movable feasts/fasts whose dates are relative to the date of Easter
var RelativeToEasterDays = []RelativeToEaster{
	{"Septuagesima Sunday", -63},
	{"Sexagesima Sunday", -56},
	{"Quinguagesima Sunday", -49},
	{"Ash Wednesday", -46},
	{"Ember Wednesday (Lent)", -39},
	{"Ember Friday (Lent)", -37},
	{"Ember Saturday (Lent)", -36},
	{"Passion Sunday", -14},
	{"Palm Sunday", -7},
	{"Spy Wednesday", -4},
	{"Holy Thursday", -3},
	{"Good Friday", -2},
	{"Holy Saturday", -1},
	{"Ascension", 39},
	{"Pentecost", 49},
	{"Corpus Christi", 60},
}

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

// relativeToEaster calculates and returns the date of a given feast/fast for a given year by applying an offset to the date of Easter for the same year.
func relativeToEaster(year int, name string) time.Time {
	for _, r := range RelativeToEasterDays {
		if r.Name == name {
			return Easter(year).AddDate(0, 0, r.Offset)
		}
	}
	panic("Unknown feast/fast: " + name)
}

// AshWednesday calculates the date of Ash Wednesday for a given year
func AshWednesday(year int) time.Time { return relativeToEaster(year, "Ash Wednesday") }

// PalmSunday calculates the date of Palm Sunday for a given year
func PalmSunday(year int) time.Time { return relativeToEaster(year, "Palm Sunday") }

// SpyWednesday calculates the date of Spy Wednesday for a given year
func SpyWednesday(year int) time.Time { return relativeToEaster(year, "Spy Wednesday") }

// HolyThursday calculates the date of Holy Thursday for a given year
func HolyThursday(year int) time.Time { return relativeToEaster(year, "Holy Thursday") }

// GoodFriday calculates the date of Good Friday for a given year
func GoodFriday(year int) time.Time { return relativeToEaster(year, "Good Friday") }

// HolySaturday calculates the date of Holy Saturday for a given year
func HolySaturday(year int) time.Time { return relativeToEaster(year, "Holy Saturday") }

// Ascension calculates the date of Ascension for a given year
func Ascension(year int) time.Time { return relativeToEaster(year, "Ascension") }

// Pentecost calculates the date of Pentecost for a given year
func Pentecost(year int) time.Time { return relativeToEaster(year, "Pentecost") }

// CorpusChristi calculates the date of CorpusChristi for a given year
func CorpusChristi(year int) time.Time { return relativeToEaster(year, "Corpus Christi") }
