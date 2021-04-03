package calculator

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {
	config := Config{
		MonthlyInvest:     MXN(100),
		StartDate:         onlyTime(time.RFC3339, "2020-01-01T00:00:00Z"),
		EndDateInvestment: onlyTime(time.RFC3339, "2022-01-01T00:00:00Z"),
		EndDate:           onlyTime(time.RFC3339, "2023-01-01T00:00:00Z"),
		Interest:          10,
		Tax:               0,
	}
	res := Calculate(config)
	// Two year investments:
	// First anual return:
	far := calcReturn(100*12, 10)
	// Second anual return reinvestment:
	reinvfar := calcReturn(uint32(far), 10)
	// Second anual return:
	scnd := calcReturn(100*12, 10)

	// Final year, not reinvesting:
	cash := MXN(100 * 12)

	fmt.Printf("far: %v\t reinvfar: %v\t scnd: %v\t cash: %v\n", far, reinvfar, scnd, cash)

	expected := MXN(cash + reinvfar + scnd)
	assert.Equal(t, expected, res)
	fmt.Printf("res: %v\n", res)
}

func TestCalculateReturn(t *testing.T) {
	inv := Investment{
		Amount:   110,
		Interest: 10,
	}
	inv.CalculateReturn()
	assert.Equal(t, inv.Return, MXN(121))
}

func TestCalculateReturnAfterTax(t *testing.T) {
	inv := Investment{
		Amount:   150,
		Interest: 10,
	}
	inv.CalculateReturnAfterTax()
	assert.Equal(t, inv.ReturnAfterTax, MXN(165))
}

// CalculateReturn calculates return investment amount
func calcReturn(amount, interest uint32) MXN {
	// Change to cents
	cents := amount * 100
	r := cents + (cents/100)*uint32(interest)
	return MXN(r / 100)
}
