package calculator

import (
	"time"
)

// doing a search using linear search algorithm, nothing fancy
func searchOutdatedInvestments(investments []Investment, date time.Time) []int {
	idxs := make([]int, 0)
	for i, inv := range investments {
		if date.After(inv.EndDate) || date.Equal(inv.EndDate) {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func removeSeveralU(s []Investment, idxs []int) []Investment {
	for i, idx := range idxs {
		r := i + 1
		s[idx] = s[len(s)-r]
	}
	return s[:len(s)-len(idxs)]
}
