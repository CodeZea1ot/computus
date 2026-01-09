package computus

import (
	"testing"
	"time"

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

// TestSundayLettersNonLeap tests SundayLetters for regular years
func TestSundayLettersNonLeap(t *testing.T) {
	tests := []struct {
		year   int
		letter string
	}{
		{2023, "A"}, // Jan 1, 2023 is Sunday
		{2021, "C"}, // Jan 1, 2021 is Friday
		{2019, "F"}, // Jan 1, 2019 is Tuesday
	}

	for _, tt := range tests {
		letter, second := SundayLetters(tt.year)
		if letter != tt.letter {
			t.Errorf("SundayLetters(%d) = %s, want %s", tt.year, letter, tt.letter)
		}
		if second != "" {
			t.Errorf("SundayLetters(%d) second letter = %s, want empty", tt.year, second)
		}
	}
}

// TestSundayLettersLeapYear tests SundayLetters for leap years
func TestSundayLettersLeapYear(t *testing.T) {
	tests := []struct {
		year         int
		firstLetter  string
		secondLetter string
	}{
		{2024, "G", "F"}, // Jan 1, 2024 Monday
		{2000, "B", "A"}, // Jan 1, 2000 Saturday
		{1996, "G", "F"}, // Jan 1, 1996 Monday
	}

	for _, tt := range tests {
		first, second := SundayLetters(tt.year)
		if first != tt.firstLetter {
			t.Errorf("SundayLetters(%d) first = %s, want %s", tt.year, first, tt.firstLetter)
		}
		if second != tt.secondLetter {
			t.Errorf("SundayLetters(%d) second = %s, want %s", tt.year, second, tt.secondLetter)
		}
	}
}

// TestIsLeapYear tests isLeapYear for various century and regular years
func TestIsLeapYear(t *testing.T) {
	tests := []struct {
		year int
		leap bool
	}{
		{2020, true},  // divisible by 4
		{2021, false}, // not divisible by 4
		{1900, false}, // divisible by 100 but not 400
		{2000, true},  // divisible by 400
		{2100, false}, // divisible by 100 but not 400
		{2400, true},  // divisible by 400
	}

	for _, tt := range tests {
		got := isLeapYear(tt.year)
		if got != tt.leap {
			t.Errorf("isLeapYear(%d) = %v, want %v", tt.year, got, tt.leap)
		}
	}
}

// TestSundayLettersRange performs a sanity check on SundayLetters over many years
func TestSundayLettersRange(t *testing.T) {
	letters := "ABCDEFG"
	for year := 1583; year <= 3000; year++ {
		first, second := SundayLetters(year)

		// Check first letter corresponds to the weekday of Jan 1
		jan1 := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		weekday := int(jan1.Weekday()) // Sunday=0
		expectedFirst := sundayLetterFromWeekday(weekday)
		if first != expectedFirst {
			t.Errorf("Year %d: first letter = %s, want %s", year, first, expectedFirst)
		}

		// Check second letter for leap years
		if isLeapYear(year) {
			expectedSecond := letters[(7-weekday-1+7)%7 : (7-weekday-1+7)%7+1]
			if second != expectedSecond {
				t.Errorf("Year %d: second letter = %s, want %s", year, second, expectedSecond)
			}
		} else {
			if second != "" {
				t.Errorf("Year %d: second letter = %s, want empty", year, second)
			}
		}
	}
}

// sundayLetterFromWeekday converts weekday (0=Sunday..6=Saturday) to Sunday letter A..G
func sundayLetterFromWeekday(weekday int) string {
	letters := "ABCDEFG"
	// Sunday = 0 => A, Saturday=6 => B (wrapping with formula)
	return letters[(7-weekday)%7 : (7-weekday)%7+1]
}
