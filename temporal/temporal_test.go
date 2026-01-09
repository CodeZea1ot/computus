package temporal

import (
	"testing"

	"github.com/CodeZea1ot/computus"
)

func TestRelativeToEasterKnownFeasts(t *testing.T) {
	feasts := []string{
		"Ash Wednesday",
		"Palm Sunday",
		"Good Friday",
		"Ascension",
		"Pentecost",
		EmberWedLent,
		EmberFriPent,
	}

	for _, name := range feasts {
		got := mustRelativeToEaster(2026, name)
		if got.IsZero() {
			t.Errorf("mustRelativeToEaster returned zero for %s", name)
		}

		// Optional: verify that the offset matches Easter
		easter := computus.Easter(2026)
		var expectedOffset int
		for _, r := range RelativeToEasterDays {
			if r.Name == name {
				expectedOffset = r.Offset
				break
			}
		}
		diff := int(got.Sub(easter).Hours() / 24)
		if diff != expectedOffset {
			t.Errorf("%s: got %d days from Easter, want %d", name, diff, expectedOffset)
		}
	}
}

func TestMustRelativeToEasterInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for unknown feast")
		}
	}()
	mustRelativeToEaster(2026, "Nonexistent Feast")
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
		easter := computus.Easter(year)

		diff := int(dayToCheck.Sub(easter).Hours() / 24)
		if diff != offset {
			t.Fatalf("%s(%d) is %d days from Easter, want %d", name, year, diff, offset)
		}
	}
}

// TestRelativeToEasterInRange tests the validity of all feasts/fasts whose assigned date is relative to Easter
func TestRelativeToEasterInRange(t *testing.T) {
	for _, r := range RelativeToEasterDays {
		checkDaysFromEaster(t, r.Name)
	}
}
