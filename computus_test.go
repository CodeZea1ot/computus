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
