package service

import (
	"encoding/json"
	"fmt"

	"git.zx-tech.net/pengfeng/facebook/http"
	"git.zx-tech.net/pengfeng/facebook/marketing/model"
)

func GetBusinessUser(accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "id,name,business,role,title,email,pending_email"
	uri := fmt.Sprintf("v11.0/me/business_users")
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	} else {
		var res model.BusinessUsers
		err = json.Unmarshal(data, &res)
		if err != nil {
			fmt.Println("json error", err.Error())
		} else {
			data1, _ := json.Marshal(res)
			fmt.Println(string(data1))
		}
	}

	return
}

func GetBusiness(accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "id,name,business,role,email,pending_email"
	uri := fmt.Sprintf("v11.0/me/businesses")
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	} else {
		var res model.Businesses
		err = json.Unmarshal(data, &res)
		if err != nil {
			fmt.Println("json error", err.Error())
		} else {
			data1, _ := json.Marshal(res)
			fmt.Println(string(data1))
		}
	}

	return
}

func GetBusinessAdCount(userId, accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	uri := fmt.Sprintf("v11.0/%v/assigned_ad_accounts", userId)
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	} else {
		var res model.BusinessAssignedAdAccountMeta
		err = json.Unmarshal(data, &res)
		if err != nil {
			fmt.Println("json error", err.Error())
		} else {
			data1, _ := json.Marshal(res)
			fmt.Println(string(data1))
		}
	}

	return
}

//自有帐号
func GetBusinessOwnedAdCount(businessID, accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "name,business,id,balance,disable_reason,account_status,amount_spent,timezone_offset_hours_utc,spend_cap"
	uri := fmt.Sprintf("v11.0/%v/owned_ad_accounts", businessID)
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	} else {
		var res model.BusinessAssignedAdAccountMeta
		err = json.Unmarshal(data, &res)
		if err != nil {
			fmt.Println("json error", err.Error())
		} else {
			data1, _ := json.Marshal(res)
			fmt.Println(string(data1))
		}
	}

	return
}

//代理帐号
func GetBusinessClientAdCount(businessID, accountId, token string) {
	mp := make(map[string]string)
	mp["access_token"] = token
	mp["fields"] = "account_id,name,business,id,balance,disable_reason,account_status,amount_spent,timezone_offset_hours_utc,spend_cap,age,currency"
	uri := fmt.Sprintf("v11.0/%v/client_ad_accounts", businessID)
	data, err := http.Action("GET", uri, "", mp, nil)
	if err != nil {
		err = fmt.Errorf("获取数据失败：%v", err.Error())
		return
	} else {
		var res model.BusinessAssignedAdAccountMeta
		err = json.Unmarshal(data, &res)
		if err != nil {
			fmt.Println("json error", err.Error())
		} else {
			data1, _ := json.Marshal(res)
			fmt.Println(string(data1))
		}
	}

	return
}
