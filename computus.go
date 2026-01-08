package computus

import "time"

// RelativeToEaster represents a moveable feast or fast whose date is relative to Easter
type RelativeToEaster struct {
	Name   string // Name of the feast/fast
	Offset int    // Days relative to Easter (negative = before, positive = after)
}

const (
	EmberWedLent = "Ember Wednesday (Lent)"
	EmberFriLent = "Ember Friday (Lent)"
	EmberSatLent = "Ember Saturday (Lent)"
	EmberWedPent = "Ember Wednesday (Pentecost)"
	EmberFriPent = "Ember Friday (Pentecost)"
	EmberSatPent = "Ember Saturday (Pentecost)"
)

// RelativeToEasterDays represents a collection of movable feasts/fasts whose dates are relative to the date of Easter
var RelativeToEasterDays = []RelativeToEaster{
	{"Septuagesima Sunday", -63},
	{"Sexagesima Sunday", -56},
	{"Quinguagesima Sunday", -49},
	{"Ash Wednesday", -46},
	{EmberWedLent, -39},
	{EmberFriLent, -37},
	{EmberSatLent, -36},
	{"Passion Sunday", -14},
	{"Palm Sunday", -7},
	{"Spy Wednesday", -4},
	{"Holy Thursday", -3},
	{"Good Friday", -2},
	{"Holy Saturday", -1},
	{"Easter Monday", 1},
	{"Easter Tuesday", 2},
	{"The Octave of Easter (Low Sunday)", 7},
	{"Ascension", 39},
	{"Pentecost", 49},
	{EmberWedPent, 52},
	{EmberFriPent, 54},
	{EmberSatPent, 55},
	{"Trinity Sunday", 56},
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

// relativeToEaster returns the date of a movable feast or fast for the given year,
// calculated by applying the feast's offset relative to Easter Sunday.
//
// The boolean return value reports whether the provided name matches a known
// feast or fast defined in RelativeToEasterDays. If no match is found, the
// returned time value is the zero time and the boolean will be false.
//
// This function does not panic and is safe to use with untrusted input.
func relativeToEaster(year int, name string) (time.Time, bool) {
	for _, r := range RelativeToEasterDays {
		if r.Name == name {
			return Easter(year).AddDate(0, 0, r.Offset), true
		}
	}
	return time.Time{}, false
}

// mustRelativeToEaster returns the date of a movable feast or fast for the given year,
// calculated by applying the feast's offset relative to Easter Sunday.
//
// It panics if the provided name does not correspond to a known feast or fast.
// This function is intended for internal use where the feast name is a known
// constant and its presence in RelativeToEasterDays is an invariant.
func mustRelativeToEaster(year int, name string) time.Time {
	if d, ok := relativeToEaster(year, name); ok {
		return d
	}
	panic("computus: unknown feast/fast: " + name)
}

// AshWednesday calculates the date of Ash Wednesday for a given year
func AshWednesday(year int) time.Time { return mustRelativeToEaster(year, "Ash Wednesday") }

// PalmSunday calculates the date of Palm Sunday for a given year
func PalmSunday(year int) time.Time { return mustRelativeToEaster(year, "Palm Sunday") }

// SpyWednesday calculates the date of Spy Wednesday for a given year
func SpyWednesday(year int) time.Time { return mustRelativeToEaster(year, "Spy Wednesday") }

// HolyThursday calculates the date of Holy Thursday for a given year
func HolyThursday(year int) time.Time { return mustRelativeToEaster(year, "Holy Thursday") }

// GoodFriday calculates the date of Good Friday for a given year
func GoodFriday(year int) time.Time { return mustRelativeToEaster(year, "Good Friday") }

// HolySaturday calculates the date of Holy Saturday for a given year
func HolySaturday(year int) time.Time { return mustRelativeToEaster(year, "Holy Saturday") }

// Ascension calculates the date of Ascension for a given year
func Ascension(year int) time.Time { return mustRelativeToEaster(year, "Ascension") }

// Pentecost calculates the date of Pentecost for a given year
func Pentecost(year int) time.Time { return mustRelativeToEaster(year, "Pentecost") }

// CorpusChristi calculates the date of CorpusChristi for a given year
func CorpusChristi(year int) time.Time { return mustRelativeToEaster(year, "Corpus Christi") }

// OctaveOfEaster calculates the date of The Octave of Easter (Low Sunday) for a given year
func OctaveOfEaster(year int) time.Time {
	return mustRelativeToEaster(year, "The Octave of Easter (Low Sunday)")
}

// LowSunday is a wrapper around OctaveOfEaster. It calculates the date of The Octave of Easter (Low Sunday) for a given year
func LowSunday(year int) time.Time {
	return OctaveOfEaster(year)
}

// EasterMonday calculates the date of Easter Monday for a given year
func EasterMonday(year int) time.Time { return mustRelativeToEaster(year, "Easter Monday") }

// EasterTuesday calculates the date of Easter Tuesday for a given year
func EasterTuesday(year int) time.Time { return mustRelativeToEaster(year, "Easter Tuesday") }

// TrinitySunday calculates the date of Trinity Sunday for a given year
func TrinitySunday(year int) time.Time { return mustRelativeToEaster(year, "Trinity Sunday") }

// EmberWednesdayLent calculates the date of Ember Wednesday in Lent
func EmberWednesdayLent(year int) time.Time {
	return mustRelativeToEaster(year, EmberWedLent)
}

// EmberFridayLent calculates the date of Ember Friday in Lent
func EmberFridayLent(year int) time.Time {
	return mustRelativeToEaster(year, EmberFriLent)
}

// EmberSaturdayLent calculates the date of Ember Saturday in Lent
func EmberSaturdayLent(year int) time.Time {
	return mustRelativeToEaster(year, EmberSatLent)
}

// EmberWednesdayPentecost calculates the date of Ember Wednesday after Pentecost
func EmberWednesdayPentecost(year int) time.Time {
	return mustRelativeToEaster(year, EmberWedPent)
}

// EmberFridayPentecost calculates the date of Ember Friday after Pentecost
func EmberFridayPentecost(year int) time.Time {
	return mustRelativeToEaster(year, EmberFriPent)
}

// EmberSaturdayPentecost calculates the date of Ember Saturday after Pentecost
func EmberSaturdayPentecost(year int) time.Time {
	return mustRelativeToEaster(year, EmberSatPent)
}
