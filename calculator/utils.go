package calculator

import "time"

// doing a search using linear search algorithm, nothing fancy
func searchOutdatedInvestments(investments []Investment, date time.Time) []int {
	idxs := make([]int, 0)
	for i, inv := range investments {
		if inv.EndDate.After(date) {
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
	return s[:len(s)-len(idxs)] // TODO: review if len(idxs) is the correct value or: len(idxs) -/+ 1
}
