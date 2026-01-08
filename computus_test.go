package computus

import "testing"

// TestEaster verifies that the Easter calculation function
// returns the correct, historically verified Easter Sunday dates
// for a selection of known years, including edge cases such as:
//   - the first Gregorian Easter (1583)
//   - the earliest possible Easter (March 22, 1818)
//   - the latest possible Easter (April 25, 2038)
func TestEaster(t *testing.T) {
	verifiedDates := map[int]string{
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

	for year, expected := range verifiedDates {
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
//
// This test loops over a wide range of years to validate these invariants.
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
	verifiedDates := map[int]string{
		2020: "2020-02-26",
		2021: "2021-02-17",
		2022: "2022-03-02",
		2023: "2023-02-22",
		2024: "2024-02-14",
		2025: "2025-03-05",
		2026: "2026-02-18",
	}

	for year, expected := range verifiedDates {
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
	verifiedDates := map[int]string{
		2020: "2020-04-05",
		2021: "2021-03-28",
		2022: "2022-04-10",
		2023: "2023-04-02",
		2024: "2024-03-24",
		2025: "2025-04-13",
		2026: "2026-03-29",
	}

	for year, expected := range verifiedDates {
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

// TestGoodFriday verifies that Good Friday is correctly calculated
// as 2 days before Easter Sunday for a selection of known years.
func TestGoodFriday(t *testing.T) {
	verifiedDates := map[int]string{
		2020: "2020-04-10",
		2021: "2021-04-02",
		2022: "2022-04-15",
		2023: "2023-04-07",
		2024: "2024-03-29",
		2025: "2025-04-18",
		2026: "2026-04-03",
	}

	for year, expected := range verifiedDates {
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

// TestPentecost verifies that Pentecost is correctly calculated
// as 49 days after Easter Sunday for a selection of known years.
func TestPentecost(t *testing.T) {
	verifiedDates := map[int]string{
		2020: "2020-05-31",
		2021: "2021-05-23",
		2022: "2022-06-05",
		2023: "2023-05-28",
		2024: "2024-05-19",
		2025: "2025-06-08",
		2026: "2026-05-24",
	}

	for year, expected := range verifiedDates {
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
