package calculator

import (
	"fmt"
	"time"
)

type percentage uint8

// Config configuration calculate compound interest
type Config struct {
	MonthlyInvest     MXN        `yaml:"monthly_invest"`
	StartDate         time.Time  `yaml:"start_date"`
	EndDate           time.Time  `yaml:"end_date"`
	EndDateInvestment time.Time  `yaml:"end_date_investment"`
	Interest          percentage `yaml:"interest"`
	Tax               percentage `yaml:"tax"`
}

// MXN represents Mexican Peso, in cents
type MXN int64

// Investment is an investment
type Investment struct {
	StartDate      time.Time
	EndDate        time.Time
	Amount         MXN
	Interest       percentage
	Return         MXN
	ReturnAfterTax MXN
}

// CalculateReturn calculates return investment amount
func (i *Investment) CalculateReturn() {
	// Change to cents
	cents := uint32(i.Amount * 100)
	r := cents + (cents/100)*uint32(i.Interest)
	i.Return = MXN(r / 100)
}

// CalculateReturnAfterTax calculate returns investment after paying tax
func (i *Investment) CalculateReturnAfterTax() {
	i.CalculateReturn()
	i.ReturnAfterTax = i.Return // TODO: Calculate correctly
}

// Calculate calculates compound interest and returns total sum
func Calculate(config Config) MXN {
	portfolio := Portfolio{
		Investments:     make([]Investment, 0),
		PastInvestments: make([]Investment, 0),
	}
	currentDate := config.StartDate
	for ; currentDate.Before(config.EndDateInvestment); currentDate = currentDate.AddDate(0, 1, 0) {
		fmt.Printf("currentDate: %s\n", currentDate)
		oldInvestments := portfolio.Refresh(currentDate)
		for _, oldInv := range oldInvestments {
			portfolio.Reinvest(oldInv)
		}

		newInvestment := Investment{
			StartDate: currentDate,
			EndDate:   currentDate.AddDate(1, 0, 0),
			Amount:    config.MonthlyInvest,
			Interest:  config.Interest,
		}
		portfolio.Invest(newInvestment)
		fmt.Printf("==================================================\n")
	}
	for ; currentDate.Before(config.EndDate); currentDate = currentDate.AddDate(0, 1, 0) {
		portfolio.InsertCash(config.MonthlyInvest)
		fmt.Printf("currentDate: %s\n", currentDate)
		fmt.Printf("==================================================\n")
	}
	return portfolio.NetWorth
}
