package sanctoral

import (
	"testing"

	"github.com/CodeZea1ot/computus/rubric"
)

func TestFixedDayCreation(t *testing.T) {
	fd := FixedDay{
		Name:     "St. Example",
		Month:    4,
		Day:      25,
		Rank:     rubric.GreaterDouble,
		Optional: false,
	}

	if fd.Name != "St. Example" || fd.Month != 4 || fd.Day != 25 {
		t.Errorf("FixedDay fields not set correctly: %+v", fd)
	}

	if fd.Rank != rubric.GreaterDouble {
		t.Errorf("Rank incorrect: got %v, want %v", fd.Rank, rubric.GreaterDouble)
	}

	if fd.Optional != false {
		t.Errorf("Optional field incorrect: got %v, want false", fd.Optional)
	}
}
