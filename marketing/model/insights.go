package model

import "git.zx-tech.net/pengfeng/facebook/model"

type Insights struct {
	Data   []Insight `json:"data"`
	Paging model.Paging    `json:"paging"`
}

// Insight 广告的分析
type Insight struct {
	AccountId                  string                       `json:"account_id"`
	AccountName                string                       `json:"account_name"`
	AdId                       string                       `json:"ad_id"`
	AdName                     string                       `json:"ad_name"`
	AdsetId                    string                       `json:"adset_id"`
	AdsetName                  string                       `json:"adset_name"`
	CampaignId                 string                       `json:"campaign_id"`
	CampaignName               string                       `json:"campaign_name"`
	BuyingType                 string                       `json:"buying_type"`
	SocialSpend                float64                      `json:"social_spend,string"`
	AccountCurrency            string                       `json:"account_currency"`
	Spend                      float64                      `json:"spend,string"`
	Cpc                        float64                      `json:"cpc,string"`
	Cpm                        float64                      `json:"cpm,string"`
	Ctr                        float64                      `json:"ctr,string"` //fb已经乘以100  按百分比显示
	Impressions                int64                        `json:"impressions,string"`
	Clicks                     int64                        `json:"clicks,string"`
	Reach                      int64                        `json:"reach,string"`
	ConversionRateRanking      string                       `json:"conversion_rate_ranking"`
	EngagementRateRanking      string                       `json:"engagement_rate_ranking"`
	QualityRanking             string                       `json:"quality_ranking"`
	DateStart                  string                       `json:"date_start"`
	DateStop                   string                       `json:"date_stop"`
	VideoAvgTimeWatchedActions []VideoAvgTimeWatchedActions `json:"video_avg_time_watched_actions"`
	Actions                    []Actions                    `json:"actions"`
	ActionValues               []ActionValues               `json:"action_values"`
}

type VideoAvgTimeWatchedActions struct {
	ActionType string `json:"action_type"`
	Value      int64  `json:"value,string"`
}
type Actions struct {
	ActionType string `json:"action_type"`
	Value      int64  `json:"value,string"`
}
type ActionValues struct {
	ActionType string  `json:"action_type"`
	Value      float64 `json:"value,string"`
}
