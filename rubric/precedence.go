package rubric

// Rank represents the liturgical rank of a feast as defined in the pre-1962 rubrics and related sources
type Rank string

const (
	Double        Rank = "Double"
	GreaterDouble Rank = "Greater Double"
	Semidouble    Rank = "Semidouble"
	Simple        Rank = "Simple"
)
