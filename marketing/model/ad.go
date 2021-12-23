package model

import "git.zx-tech.net/pengfeng/facebook/model"

type Ad struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	Status          string      `json:"status"`
	EffectiveStatus string      `json:"effective_status"`
	AdCreatives     AdCreatives   `json:"adcreatives"`
	Creative     	AdCreative    `json:"creative"`
	Campaign     	Campaign     `json:"Campaign"`
	Insights        Insights    `json:"insights"`
}

//Ads 广告-列表
type Ads struct {
	Data   []Ad   `json:"data"`
	Paging model.Paging `json:"paging"`
}