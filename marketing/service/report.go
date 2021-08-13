package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"git.zx-tech.net/pengfeng/facebook/http"
	"git.zx-tech.net/pengfeng/facebook/marketing/model"
)

var (
	fieldCampaignInsights = `insights{actions,action_values,spend,cpc,cpm,ctr,impressions,clicks,video_avg_time_watched_actions,reach},`
	fieldSetInsights      = `insights{actions,action_values,spend,cpc,cpm,ctr,impressions,clicks,video_avg_time_watched_actions,reach},`
	fieldAdInsights       = `insights{actions,action_values,spend,cpc,cpm,ctr,impressions,clicks,video_avg_time_watched_actions,reach,conversion_rate_ranking,engagement_rate_ranking,quality_ranking},`

	fieldCampaign = `name,status,effective_status,daily_budget,lifetime_budget,bid_strategy`
	fieldSet      = `name,id,daily_budget,bid_amount,status,effective_status,lifetime_budget,bid_strategy,destination_type,optimization_goal,bid_constraints,targeting,start_time,end_time,promoted_object`
	fieldAd       = `name,id,status,effective_status,adcreatives`
	fieldCreative = `id,effective_object_story_id`

	fieldAll = `adsets{` + fieldSet + `,ads{` + fieldAd + `{` + fieldCreative + `}}},` + fieldCampaign //fieldCampaign//
)

func GetReport() {
	fmt.Println("开始测试！")
	//--------从数据库获取的数据
	var name string
	var timeZone int
	var token string
	var accountId string
	//--------从数据库获取数据结束
	{
		accountId = ""
		timeZone = -8
		token = ``
		name = ""
	}

	location := time.FixedZone(name, int(timeZone*3600))
	fileds := strings.Replace(fieldAll, "{START}", time.Now().AddDate(0, 0, -60).In(location).Format("2006-01-02"), -1)
	fileds = strings.Replace(fileds, "{END}", time.Now().In(location).Format("2006-01-02"), -1)

	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = fileds
	mp["limit"] = "1000"
	uri := fmt.Sprintf("v11.0/act_%v/campaigns", accountId)
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	var res model.AdSets
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("JSON转化失败：", err.Error())
		return
	}
	for _, v1 := range res.Data {
		if v1.Ads.Data != nil {
			for _, v2 := range v1.Ads.Data {
				if v2.AdCreatives.Data != nil {
					for _, v3 := range v2.AdCreatives.Data {
						fmt.Println("ad creatives id", v3.ID)
					}
				}
			}
		}
	}
}

type ranges struct {
	Since string `json:"since"`
	Until string `json:"until"`
}

func GetReport1() {
	fmt.Println("开始测试！")
	//--------从数据库获取的数据
	var name string
	var timeZone int
	var token string
	var accountId string
	//--------从数据库获取数据结束
	{
		accountId = ""
		timeZone = -8
		token = ``
		name = ""
	}

	location := time.FixedZone(name, int(timeZone*3600))
	fileds := strings.Replace(fieldAll, "{START}", time.Now().AddDate(0, 0, -10).In(location).Format("2006-01-02"), -1)
	fileds = strings.Replace(fileds, "{END}", time.Now().In(location).Format("2006-01-02"), -1)
	var r []ranges
	{
		for i := 0; i < 100; i++ {
			r = append(r, ranges{
				Since: time.Now().AddDate(0, 0, -i).In(location).Format("2006-01-02"),
				Until: time.Now().AddDate(0, 0, -i).In(location).Format("2006-01-02"),
			})
		}
	}
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "account_id,campaign_id,campaign_name,spend,social_spend,account_currency,date_start,date_stop,dda_results"
	mp["date_preset"] = "maximum"
	rst, _ := json.Marshal(r)
	mp["time_ranges"] = string(rst)
	mp["level"] = "campaign"
	uri := fmt.Sprintf("v11.0/act_%v/insights", accountId)
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		fmt.Println("获取数据失败：", err.Error())
		return
	}
	var res model.Insights
	err = json.Unmarshal(data, &res)
	if err != nil {
		fmt.Println("JSON转化失败：", err.Error())
		return
	}
	for _, v := range res.Data {
		uri0 := fmt.Sprintf(`v11.0/%v?fields=id,account_id,adlabels,promoted_object,recommendations,ads{`+fieldAd+`{`+fieldCreative+`}}&access_token=%v`, v.CampaignId, token)
		http.Action("GET", uri0, "", nil, nil)
	}
	fmt.Println("测试获取数据结束！", len(res.Data))
	//for _,v:=range res.Data{
	//	fmt.Println(v.CampaignName,v.Spend,v.DateStart)
	//}
}

func GetReportInsights(accountId, token, name string, timeZone, day int) (list []model.Insight, err error) {
	if day < 1 {
		return
	}
	location := time.FixedZone(name, int(timeZone*3600))
	var r []ranges
	{
		for i := 0; i < day; i++ {
			r = append(r, ranges{
				Since: time.Now().AddDate(0, 0, -i).In(location).Format("2006-01-02"),
				Until: time.Now().AddDate(0, 0, -i).In(location).Format("2006-01-02"),
			})
		}
	}
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "account_id,campaign_id,campaign_name,spend,social_spend,account_currency,date_start,date_stop"
	mp["date_preset"] = "maximum"
	rst, _ := json.Marshal(r)
	mp["time_ranges"] = string(rst)
	mp["level"] = "campaign"
	uri := fmt.Sprintf("v11.0/act_%v/insights", accountId)
	data0, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	}
	if data0 == nil {
		return
	}
	var res0 model.Insights
	err = json.Unmarshal(data0, &res0)
	if err != nil {
		err = fmt.Errorf("JSON转化失败：%v", err.Error())
		return
	}
	list = res0.Data
	next := res0.Paging.Next
	if next != "" {
		for {
			if next == "" {
				break
			}
			var res1 model.Insights
			data1, err1 := http.Action("GET", next, "", nil, nil)
			if err1 != nil {
				err = fmt.Errorf("获取分页数据失败：%v", err1.Error())
				return
			}
			if data1 == nil {
				break
			}
			err1 = json.Unmarshal(data1, &res1)
			if err1 != nil {
				err = fmt.Errorf("JSON转化失败：%v", err1.Error())
				return
			}
			list = append(list, res1.Data...)
			next = res1.Paging.Next
		}
	}
	return
}
func GetReportHost(accountId, token string) (host string, err error) {
	mp := make(map[string]string)
	mp["access_token"] = token
	uri := fmt.Sprintf("v11.0/act_%v/adcreatives?access_token=%v&fields=effective_object_story_id,adlabels,object_story_id,object_story_spec&limit=500", accountId, token)
	data0, err := http.Action("GET", uri, "", nil, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	}
	if data0 == nil {
		return
	}
	var res0 model.AdCreatives
	err = json.Unmarshal(data0, &res0)
	if err != nil {
		err = fmt.Errorf("JSON转化失败：%v", err.Error())
		return
	}
	list := res0.Data
	next := res0.Paging.Next
	if next != "" {
		for {
			if next == "" {
				break
			}
			var res1 model.AdCreatives
			data1, err1 := http.Action("GET", next, "", nil, nil)
			if err1 != nil {
				err = fmt.Errorf("获取分页数据失败：%v", err1.Error())
				return
			}
			if data1 == nil {
				break
			}
			err1 = json.Unmarshal(data1, &res1)
			if err1 != nil {
				err = fmt.Errorf("JSON转化失败：%v", err1.Error())
				return
			}
			list = append(list, res1.Data...)
			next = res1.Paging.Next
		}
	}
	for _, cr := range list {
		if cr.ObjectStoryID == "" {
			link := ""
			objectStorySpec := cr.ObjectStorySpec
			if objectStorySpec.LinkData.Link != "" {
				link = objectStorySpec.LinkData.Link
			} else if objectStorySpec.VideoData.CallToAction.Value.Link != "" {
				link = objectStorySpec.VideoData.CallToAction.Value.Link
			} else if objectStorySpec.PhotoData.CallToAction.Value.Link != "" {
				link = objectStorySpec.PhotoData.CallToAction.Value.Link
			}
			if link != "" {
				host = http.GetDomain(link)
				if host != "" && !strings.Contains(host, "facebook.com") {
					return
				}
			}
		}
	}
	return
}
func GetAdCreatives(accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	uri := fmt.Sprintf("v11.0/act_%v/adcreatives?access_token=%v&fields=effective_object_story_id,adlabels,asset_feed_spec,object_story_id,object_story_spec&limit=300", accountId, token)
	data, err := http.ActionAll("GET", uri, "", nil, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	}
	if data == nil {
		return
	}
	var res model.AdCreatives
	err = json.Unmarshal(data, &res)
	if err != nil {
		err = fmt.Errorf("JSON转化失败：%v", err.Error())
		return
	}
	for _, cr := range res.Data {
		if cr.ObjectStoryID == "" {
			link := ""
			objectStorySpec := cr.ObjectStorySpec
			if objectStorySpec.LinkData.Link != "" {
				link = objectStorySpec.LinkData.Link
			} else if objectStorySpec.VideoData.CallToAction.Value.Link != "" {
				link = objectStorySpec.VideoData.CallToAction.Value.Link
			} else if objectStorySpec.PhotoData.CallToAction.Value.Link != "" {
				link = objectStorySpec.PhotoData.CallToAction.Value.Link
			}
			if link != "" {
				fmt.Println(cr.EffectiveObjectStoryID, "------->", http.GetDomain(link))
			}
		}
	}
	return
}

func GetAdImages(accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "effective_object_story_id,object_story_id,object_story_spec"
	uri := fmt.Sprintf("v11.0/act_%v/adimages", accountId)
	_, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	}

	return
}
func GetAdCreativeById(accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "object_story_spec"
	uri := fmt.Sprintf("v11.0/%v", "23847910545630562")
	_, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	}

	return
}
func GetAds(accountId, token string) {
	uri := fmt.Sprintf("v11.0/%v/ads?access_token=%v&fields=id,creative,campaign&limit=5&date_preset=this_year", accountId, token)
	data, err := http.ActionAll("GET", uri, "", nil, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	}
	if data == nil {
		fmt.Println("暂无数据")
		return
	}

	return
}
