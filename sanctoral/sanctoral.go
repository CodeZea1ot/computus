package sanctoral

import "github.com/CodeZea1ot/computus/rubric"

// FixedDay represents a saint's day or other feast/fast occurring on a fixed calendar date
type FixedDay struct {
	Name     string      // Name of the feast/fast
	Month    int         // 1=January ... 12=December
	Day      int         // 1..31
	Rank     rubric.Rank // Liturgical rank (Double, Greater Double, etc.)
	Optional bool        // True if it can be omitted (optional memorial)
}
