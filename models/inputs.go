package models

import (
	"fmt"
	"log"
	"strconv"
)

type RentOrOwn struct {
	slug string
}

func (r RentOrOwn) String() string {
	return r.slug
}

var (
	Rent = RentOrOwn{"rent"}
	Own  = RentOrOwn{"own"}
)

type PlanInput struct {
	Email                string  `json:"email"`
	Income               uint    `json:"income"`
	Age                  uint    `json:"age"`
	CCDebt               uint    `json:"ccdebt"`
	CCInterest           float32 `json:"ccinterest"`
	OtherDebt            uint    `json:"odebt"`
	OtherInterest        float32 `json:"ointerest"`
	Investments          uint    `json:"itotal"`
	TaxAdvantagedPercent uint    `json:"tap"`
	EFund                uint    `json:"efund"`
	RentOrOwn            string  `json:"rentorown"`
	MonthlyHousingCost   uint    `json:"hcost"`
	Mortgage             uint    `json:"mortgage"`
	MortgageInterest     float32 `json:"minterest"`
	HomeValuation        uint    `json:"hvalue"`
}

func PromptFromPlan(plan PlanInput) string {
	promptBase := `
	You are a trusted financial advisor who is not licensed, but will give me generic wealth management advice. 
	I am going to provide you with a summary of my financial situation.
	I will then ask a series of questions afterwards on how to better improve my situation.
	Your response should have a bulleted list of action items.
	`
	homePromptAddon := ""
	if plan.RentOrOwn == Rent.slug {
		log.Println("User rents.")
		homePromptAddon = fmt.Sprintf(`I do not own a home. My rent is %s`, strconv.Itoa(int(plan.MonthlyHousingCost)))
	}
	if plan.RentOrOwn == Own.slug {
		log.Println("User owns a home.")
		homePromptAddon = fmt.Sprintf(`
		I own a home. 
		My housing cost is $%s per month
		My existing mortgage is $%s
		The interest rate on my mortgage is %s percent,
		My estimated home valuation is $%s`,
			strconv.Itoa(int(plan.MonthlyHousingCost)),
			strconv.Itoa(int(plan.Mortgage)),
			fmt.Sprintf("%v", plan.MortgageInterest),
			strconv.Itoa(int(plan.HomeValuation)),
		)
	}

	prompt := fmt.Sprintf(`
	%s
	I make $%s gross per year.
	I am %s years old.
	I have $%s in credit card debt. The average interest rate on my credit card debt is %s percent.
	I have $%s in investments.
	%s percent of my investments are in tax-advanaged accounts.
	I have $%s saved in an emergency fund.
	`,
		promptBase,
		strconv.Itoa(int(plan.Income)),
		strconv.Itoa(int(plan.Age)),
		strconv.Itoa(int(plan.CCDebt)),
		fmt.Sprintf("%v", plan.CCInterest),
		strconv.Itoa(int(plan.Investments)),
		strconv.Itoa(int(plan.TaxAdvantagedPercent)),
		strconv.Itoa(int(plan.EFund)))
	prompt += homePromptAddon
	//log.Println(prompt)
	return prompt
}
