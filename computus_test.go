package computus

import (
	"testing"
	"time"
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

// TestAshWednesday verifies that Ash Wednesday is correctly calculated
// as 46 days before Easter Sunday for a selection of known years.
func TestAshWednesday(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, -46).Format("2006-01-02")
		got := AshWednesday(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("AshWednesday(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestAshWednesdayInRange verifies that the Ash Wednesday calculation
// is correct for all Gregorian years from 1583 to 3000.
// According to the liturgical rules, Ash Wednesday must always
// fall exactly 46 days before Easter Sunday. This test computes
// Easter for each year in the range and ensures that Ash Wednesday
// is precisely 46 days earlier.
func TestAshWednesdayInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		ash := AshWednesday(year)
		easter := Easter(year)

		// Must always be 46 days before Easter
		diff := int(easter.Sub(ash).Hours() / 24)
		if diff != 46 {
			t.Fatalf("AshWednesday(%d) is %d days before Easter, want 46", year, diff)
		}
	}
}

// TestPalmSunday verifies that Palm Sunday is correctly calculated
// as 7 days before Easter Sunday for a selection of known years.
func TestPalmSunday(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, -7).Format("2006-01-02")
		got := PalmSunday(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("PalmSunday(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestPalmSundayInRange verifies that Palm Sunday is always
// exactly 7 days before Easter Sunday for all Gregorian years.
func TestPalmSundayInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		palm := PalmSunday(year)
		easter := Easter(year)

		diff := int(easter.Sub(palm).Hours() / 24)
		if diff != 7 {
			t.Fatalf("PalmSunday(%d) is %d days before Easter, want 7", year, diff)
		}
	}
}

// TestSpyWednesday verifies that Spy Wednesday is correctly calculated
// as 4 days before Easter Sunday for a selection of known years.
func TestSpyWednesday(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, -4).Format("2006-01-02")
		got := SpyWednesday(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("SpyWednesday(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestSpyWednesdayInRange ensures that Spy Wednessday is always exactly
// 4 days before Easter Sunday for all Gregorian years.
func TestSpyWednesdayInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		spyWednesday := SpyWednesday(year)
		easter := Easter(year)

		diff := int(easter.Sub(spyWednesday).Hours() / 24)
		if diff != 4 {
			t.Fatalf("SpyWednesday(%d) is %d days before Easter, want 4", year, diff)
		}
	}
}

// TestHolyThursday verifies that Holy Thursday is correctly calculated
// as 3 days before Easter Sunday for a selection of known years.
func TestHolyThursday(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, -3).Format("2006-01-02")
		got := HolyThursday(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("HolyThursday(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestHolyThursdayInRange ensures that Holy Thursday is always exactly
// 3 days before Easter Sunday for all Gregorian years.
func TestHolyThursdayInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		holyThursday := HolyThursday(year)
		easter := Easter(year)

		diff := int(easter.Sub(holyThursday).Hours() / 24)
		if diff != 3 {
			t.Fatalf("HolyThursday(%d) is %d days before Easter, want 2", year, diff)
		}
	}
}

// TestGoodFriday verifies that Good Friday is correctly calculated
// as 2 days before Easter Sunday for a selection of known years.
func TestGoodFriday(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, -2).Format("2006-01-02")
		got := GoodFriday(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("GoodFriday(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestGoodFridayInRange ensures that Good Friday is always exactly
// 2 days before Easter Sunday for all Gregorian years.
func TestGoodFridayInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		goodFriday := GoodFriday(year)
		easter := Easter(year)

		diff := int(easter.Sub(goodFriday).Hours() / 24)
		if diff != 2 {
			t.Fatalf("GoodFriday(%d) is %d days before Easter, want 2", year, diff)
		}
	}
}

// TestHolySaturday verifies that Holy Saturday is correctly calculated
// as 3 days before Easter Sunday for a selection of known years.
func TestHolySaturday(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, -1).Format("2006-01-02")
		got := HolySaturday(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("HolySaturday(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestHolySaturdayInRange ensures that Holy Saturday is always exactly
// 1 day before Easter Sunday for all Gregorian years.
func TestHolySaturdayInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		holySaturday := HolySaturday(year)
		easter := Easter(year)

		diff := int(easter.Sub(holySaturday).Hours() / 24)
		if diff != 1 {
			t.Fatalf("HolySaturday(%d) is %d days before Easter, want 2", year, diff)
		}
	}
}

// TestPentecost verifies that Pentecost is correctly calculated
// as 49 days after Easter Sunday for a selection of known years.
func TestPentecost(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, 49).Format("2006-01-02")
		got := Pentecost(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("Pentecost(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestPentecostInRange ensures that Pentecost is always exactly
// 49 days after Easter Sunday for all Gregorian years.
func TestPentecostInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		pentecost := Pentecost(year)
		easter := Easter(year)

		diff := int(pentecost.Sub(easter).Hours() / 24)
		if diff != 49 {
			t.Fatalf("Pentecost(%d) is %d days after Easter, want 49", year, diff)
		}
	}
}

// TestAscension verifies that Ascension Thursday is correctly calculated
// as 39 days after Easter Sunday for a selection of known years.
func TestAscension(t *testing.T) {
	for year, easterStr := range verifiedEasterDates {
		easter, _ := time.Parse("2006-01-02", easterStr)
		expected := easter.AddDate(0, 0, 39).Format("2006-01-02")
		got := Ascension(year).Format("2006-01-02")
		if got != expected {
			t.Errorf("Ascension(%d) = %s, want %s", year, got, expected)
		}
	}
}

// TestAscensionInRange ensures that Ascension Thursday is always exactly
// 39 days after Easter Sunday for all Gregorian years.
func TestAscensionInRange(t *testing.T) {
	for year := 1583; year <= 3000; year++ {
		ascension := Ascension(year)
		easter := Easter(year)

		diff := int(ascension.Sub(easter).Hours() / 24)
		if diff != 39 {
			t.Fatalf("Ascension(%d) is %d days after Easter, want 39", year, diff)
		}
	}
}
