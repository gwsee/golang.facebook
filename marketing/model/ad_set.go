package model

import "git.zx-tech.net/pengfeng/facebook/model"

type AdSets struct {
	Data   []AdSet      `json:"data"`
	Paging model.Paging `json:"paging"`
}

//AdSet 广告集
type AdSet struct {
	ID               string                 `json:"id"`
	Name             string                 `json:"name"`
	DailyBudget      int64                  `json:"daily_budget,string"`
	BidAmount        int64                  `json:"bid_amount"`
	LifetimeBudget   int64                  `json:"lifetime_budget,string"`
	BidStrategy      string                 `json:"bid_strategy"`
	Status           string                 `json:"status"`
	EffectiveStatus  string                 `json:"effective_status"`
	DestinationType  string                 `json:"destination_type"`
	OptimizationGoal string                 `json:"optimization_goal"`
	Targeting        map[string]interface{} `json:"targeting"`
	StartTime        string                 `json:"start_time"`
	EndTime          string                 `json:"end_time"`
	PromotedObject   PromotedObject         `json:"promoted_object"`
	Ads              Ads                    `json:"ads"`
	Insights         Insights               `json:"insights"`
}

type PromotedObject struct {
	PixelID         string `json:"pixel_id"`
	PixelName       string `json:"pixel_name,omitempty"`
	CustomEventType string `json:"custom_event_type"`
}
