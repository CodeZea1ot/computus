package computus

import (
	"testing"

	"github.com/CodeZea1ot/computus/internal/testutil"
)

// TestEaster verifies that the Easter calculation function
// returns the correct, historically verified Easter Sunday dates
// for a selection of known years, including edge cases such as:
//   - the first Gregorian Easter (April 4, 1583)
//   - the earliest possible Easter (March 22)
//   - the latest possible Easter (April 25)
func TestEaster(t *testing.T) {
	for year, expected := range testutil.VerifiedEasterDates {
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
