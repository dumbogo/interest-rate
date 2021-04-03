package calculator

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRefresh(t *testing.T) {
	invs := []Investment{
		{
			StartDate:      onlyTime(time.RFC3339, "2020-01-01T00:00:00Z"),
			EndDate:        onlyTime(time.RFC3339, "2021-01-01T00:00:00Z"),
			Amount:         500,
			Interest:       10,
			Return:         550,
			ReturnAfterTax: 550,
		},
		{
			StartDate:      onlyTime(time.RFC3339, "2020-02-01T00:00:00Z"),
			EndDate:        onlyTime(time.RFC3339, "2021-02-01T00:00:00Z"),
			Amount:         100,
			Interest:       10,
			Return:         110,
			ReturnAfterTax: 110,
		},
		{
			StartDate:      onlyTime(time.RFC3339, "2020-03-01T00:00:00Z"),
			EndDate:        onlyTime(time.RFC3339, "2021-03-01T00:00:00Z"),
			Amount:         100,
			Interest:       10,
			Return:         110,
			ReturnAfterTax: 110,
		},
	}
	expected := make([]Investment, 2)
	expectedCurrentInvestments := make([]Investment, 1)
	copy(expected, invs[:2])
	copy(expectedCurrentInvestments, invs[2:])
	testCase := Portfolio{
		Investments: invs,
	}
	actual := testCase.Refresh(onlyTime(time.RFC3339, "2021-03-01T00:00:00Z"))
	assert.Equal(t, 2, len(testCase.PastInvestments), "Length difference in PastInvestments")
	assert.Equal(t, 1, len(testCase.Investments), "Length difference Investments")
	assert.EqualValues(t, expected, actual)
	assert.EqualValues(t, expected, testCase.PastInvestments)
	assert.EqualValues(t, expectedCurrentInvestments, testCase.Investments)
}

func TestInvest(t *testing.T) {
	investment := Investment{
		StartDate:      onlyTime(time.RFC3339, "2020-03-01T00:00:00Z"),
		EndDate:        onlyTime(time.RFC3339, "2021-03-01T00:00:00Z"),
		Amount:         100,
		Interest:       10,
		Return:         110,
		ReturnAfterTax: 110,
	}
	portfolio := Portfolio{}
	portfolio.Invest(investment)
	expected := make([]Investment, 1)
	expected[0] = investment
	assert.EqualValues(t, expected, portfolio.Investments)
	assert.Equal(t, MXN(110), portfolio.NetWorth, "Net worth is not ok")
}

func TestInsertCash(t *testing.T) {
	portfolio := Portfolio{}
	portfolio.InsertCash(MXN(300))
	assert.Equal(t, MXN(300), portfolio.NetWorth)
	assert.Equal(t, MXN(300), portfolio.Cash)
}

func TestString(t *testing.T) {
	expected := `Investments: []
PastInvestments: []` +
		"\nNet worth: 0 \t Cash: 0\n"
	actual := Portfolio{}.String()
	fmt.Printf(actual)
	assert.Equal(t, expected, actual)
}

func onlyTime(format, val string) time.Time {
	time, e := time.Parse(format, val)
	check(e)
	return time
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
