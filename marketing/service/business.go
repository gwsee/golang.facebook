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
	mp["fields"] = "id,name,business,role,email,pending_email"
	uri := fmt.Sprintf("v11.0/me/business_users")
	// 返回
	// {"data":[{"id":"106914815003833","name":"Lillian Knight","business":{"id":"200643245322999","name":"\u5b81\u590f\u91d1\u8fd0\u83b1\u5546\u8d38\u6709\u9650\u516c\u53f8"},"role":"ADMIN","pending_email":"support\u0040hepeach.com"}],"paging":{"cursors":{"before":"QVFIUnd6d2JyMWMyNkJJRXZAVMTFHdG9iVURjM2kyanBTU1pYempOb3B2bXVpMzduNkdDakFfcWRaTGxmLWh1c21yVGhoOGVoQ2dhMlkxZAEpRUmlJbEdGbVRR","after":"QVFIUnd6d2JyMWMyNkJJRXZAVMTFHdG9iVURjM2kyanBTU1pYempOb3B2bXVpMzduNkdDakFfcWRaTGxmLWh1c21yVGhoOGVoQ2dhMlkxZAEpRUmlJbEdGbVRR"}}}
	// 处理后
	// {"data":[{"id":"106914815003833","name":"Lillian Knight","role":"ADMIN","title":"","email":"","finance_permission":"","first_name":"","last_name":"","pending_email":"support@hepeach.com","ip_permission":"","business":{"id":"200643245322999","name":"宁夏金运莱商贸有限公司","link":""}}],"paging":{"cursors":{"before":"QVFIUnd6d2JyMWMyNkJJRXZAVMTFHdG9iVURjM2kyanBTU1pYempOb3B2bXVpMzduNkdDakFfcWRaTGxmLWh1c21yVGhoOGVoQ2dhMlkxZAEpRUmlJbEdGbVRR","after":"QVFIUnd6d2JyMWMyNkJJRXZAVMTFHdG9iVURjM2kyanBTU1pYempOb3B2bXVpMzduNkdDakFfcWRaTGxmLWh1c21yVGhoOGVoQ2dhMlkxZAEpRUmlJbEdGbVRR"},"next":""}}
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
	// 返回
	// {"data":[{"id":"200643245322999","name":"\u5b81\u590f\u91d1\u8fd0\u83b1\u5546\u8d38\u6709\u9650\u516c\u53f8"}],"paging":{"cursors":{"before":"QVFIUjBxaDRCYVdPSXdGX3BCVndwTHVpQU1TWDdtdkdXbl95TTRUaEluQVF4QnRHWUxrMTU3dDZAldXV5ZAlNTV2FDc2plY1RwY2pTSWVOdmloOVFSRUd5NHlB","after":"QVFIUjBxaDRCYVdPSXdGX3BCVndwTHVpQU1TWDdtdkdXbl95TTRUaEluQVF4QnRHWUxrMTU3dDZAldXV5ZAlNTV2FDc2plY1RwY2pTSWVOdmloOVFSRUd5NHlB"}}}
	// 处理后
	// {"data":[{"id":"200643245322999","name":"宁夏金运莱商贸有限公司","link":""}],"paging":{"cursors":{"before":"QVFIUjBxaDRCYVdPSXdGX3BCVndwTHVpQU1TWDdtdkdXbl95TTRUaEluQVF4QnRHWUxrMTU3dDZAldXV5ZAlNTV2FDc2plY1RwY2pTSWVOdmloOVFSRUd5NHlB","after":"QVFIUjBxaDRCYVdPSXdGX3BCVndwTHVpQU1TWDdtdkdXbl95TTRUaEluQVF4QnRHWUxrMTU3dDZAldXV5ZAlNTV2FDc2plY1RwY2pTSWVOdmloOVFSRUd5NHlB"},"next":""}}
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
