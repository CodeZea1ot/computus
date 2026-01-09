package rubric

import "testing"

func TestRankConstants(t *testing.T) {
	tests := []struct {
		rank Rank
		want string
	}{
		{Double, "Double"},
		{GreaterDouble, "Greater Double"},
		{Semidouble, "Semidouble"},
		{Simple, "Simple"},
	}

	for _, tt := range tests {
		if string(tt.rank) != tt.want {
			t.Errorf("Rank %v = %q, want %q", tt.rank, string(tt.rank), tt.want)
		}
	}
}
