package calculator

import (
	"fmt"
	"time"
)

type percentage uint8

// Config configuration calculate compound interest
type Config struct {
	MonthlyInvest uint32     `yaml:"monthly_invest"`
	StartDate     time.Time  `yaml:"start_date"`
	EndDate       time.Time  `yaml:"end_date"`
	Interest      percentage `yaml:"interest"`
	Tax           percentage `yaml:"tax"`
}

// MXN represents Mexican Peso, in cents
type MXN int64

// Investment is an investment
type Investment struct {
	StartDate       time.Time
	EndDate         time.Time
	Ammount         MXN
	Interest        percentage
	Returns         MXN
	ReturnsAfterTax MXN
}

// MaxPermitedInvestments to 12,
// because StartDate and EndDate will never be more than 12 months and only one investment per month is allowed
const MaxPermitedInvestments = 12

// Portfolio is the current investments
type Portfolio struct {
	Investments     []Investment
	PastInvestments []Investment
}

// Refresh refreshes Portfolio based on the date,
// Moves Investments with EndDate > date to PastInvestments
func (P *Portfolio) Refresh(date time.Time) {
	idxsInvestmentsAfter := searchOutdatedInvestments(P.Investments, date)
	for _, idx := range idxsInvestmentsAfter {
		P.PastInvestments = append(P.Investments, P.Investments[idx])
	}
	P.Investments = removeSeveralU(P.Investments, idxsInvestmentsAfter)
}

// Reinvest reinvest all investments that passed from Investments to PastInvestments
// returns reinvested Investments
func (P *Portfolio) Reinvest(date time.Time) []Investment {
	// TODO: end this
	return nil
}

// Calculate calculates compound interest and returns total sum
// TODO: WIP
func Calculate(config Config) uint32 {
	var total uint32
	portFolio := Portfolio{
		Investments:     make([]Investment, 0),
		PastInvestments: make([]Investment, 0),
	}
	for currentDate := config.StartDate; currentDate.Before(config.EndDate); currentDate = currentDate.AddDate(0, 1, 0) {
		fmt.Printf("currentDate: %s\n", currentDate)

		portFolio.Refresh(currentDate)
		fmt.Printf("total investment with returns: %d\n", total)
	}
	return 0
}
