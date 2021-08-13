package model

type Campaign struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Status          string   `json:"status"`
	EffectiveStatus string   `json:"effective_status"`
	LifetimeBudget  int64    `json:"lifetime_budget,string"`
	DailyBudget     int64    `json:"daily_budget,string"`
	BidStrategy     string   `json:"bid_strategy"`
	Adsets          AdSets   `json:"adsets"`
	Insights        Insights `json:"insights"`
}
