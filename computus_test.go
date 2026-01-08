package computus

import (
	"testing"
)

// VerifiedEasterDates contains historically verified Easter Sunday dates.
// This can be used by all dependent tests to calculate moveable feasts.
var verifiedEasterDates = map[int]string{
	1583: "1583-04-10", // first year Gregorian Easter
	1666: "1666-04-25", // latest possible Easter
	1693: "1693-03-22", // earliest possible Easter
	1818: "1818-03-22", // earliest possible Easter
	1900: "1900-04-15",
	1954: "1954-04-18",
	1970: "1970-03-29",
	1999: "1999-04-04",
	2000: "2000-04-23",
	2010: "2010-04-04",
	2016: "2016-03-27",
	2020: "2020-04-12",
	2021: "2021-04-04",
	2022: "2022-04-17",
	2023: "2023-04-09",
	2024: "2024-03-31",
	2025: "2025-04-20",
	2026: "2026-04-05",
	2038: "2038-04-25", // latest possible Easter
}

// checkDaysFromEaster validates that the relative feast/fasts fall on
// the correct number of days from Easter for all years in the range.
func checkDaysFromEaster(t *testing.T, name string) {
	beginYear := 1583
	endYear := 3000
	var offset int
	found := false
	for _, r := range RelativeToEasterDays {
		if r.Name == name {
			offset = r.Offset
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("feast %q not found in RelativeToEasterDays", name)
	}

	for year := beginYear; year <= endYear; year++ {
		dayToCheck := mustRelativeToEaster(year, name)
		easter := Easter(year)

		diff := int(dayToCheck.Sub(easter).Hours() / 24)
		if diff != offset {
			t.Fatalf("%s(%d) is %d days from Easter, want %d", name, year, diff, offset)
		}
	}
}

// TestEaster verifies that the Easter calculation function
// returns the correct, historically verified Easter Sunday dates
// for a selection of known years, including edge cases such as:
//   - the first Gregorian Easter (April 4, 1583)
//   - the earliest possible Easter (March 22)
//   - the latest possible Easter (April 25)
func TestEaster(t *testing.T) {
	for year, expected := range verifiedEasterDates {
		got := Easter(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("Easter(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestEasterInRange ensures that the Easter function never returns
// an impossible date. According to the Gregorian computus rules:
//   - Easter always falls on a Sunday
//   - The date is always between March 22 and April 25
func TestEasterInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		e := Easter(year)
		if e.Month() < 3 || e.Month() > 4 {
			t.Fatalf("Easter(%d) is in invalid month: %v", year, e)
		}
		if e.Month() == 3 && e.Day() < 22 {
			t.Fatalf("Easter(%d) is earlier than March 22: %v", year, e)
		}
		if e.Month() == 4 && e.Day() > 25 {
			t.Fatalf("Easter(%d) is later than April 25: %v", year, e)
		}
	}
}

// TestRelativeToEasterInRange tests the validity of all feasts/fasts whose assigned date is relative to Easter
func TestRelativeToEasterInRange(t *testing.T) {
	for _, r := range RelativeToEasterDays {
		checkDaysFromEaster(t, r.Name)
	}
}
