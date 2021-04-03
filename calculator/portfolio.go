package calculator

import (
	"fmt"

	"time"
)

// Portfolio is the current investments
type Portfolio struct {
	Investments     []Investment
	PastInvestments []Investment
	NetWorth        MXN
	Cash            MXN
}

// Refresh refreshes Portfolio based on the date,
// Moves Investments with EndDate > date to PastInvestments
// returns moved Investments to PastInvestments
func (P *Portfolio) Refresh(date time.Time) []Investment {
	idxsOutdatedInvs := searchOutdatedInvestments(P.Investments, date)
	outdatedInvs := make([]Investment, 0)
	for _, idx := range idxsOutdatedInvs {
		P.PastInvestments = append(P.PastInvestments, P.Investments[idx])
		outdatedInvs = append(outdatedInvs, P.Investments[idx])
	}
	P.Investments = removeSeveralU(P.Investments, idxsOutdatedInvs)
	return outdatedInvs
}

// Invest invest new investment in portfolio,
// sums to NetWorth
func (P *Portfolio) Invest(inv Investment) {
	inv.CalculateReturnAfterTax()
	P.Investments = append(P.Investments, inv)
	P.NetWorth += inv.ReturnAfterTax
}

// InsertCash adds cash to Cash Portfolio and sums to NetWorth
func (P *Portfolio) InsertCash(m MXN) {
	P.Cash += m
	P.NetWorth += m
}

func (P Portfolio) String() string {
	return fmt.Sprintf("Investments: %+v\nPastInvestments: %+v\nNet worth: %d \t Cash: %d\n",
		P.Investments, P.PastInvestments, P.NetWorth, P.Cash,
	)
}
