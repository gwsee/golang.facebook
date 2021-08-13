package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"git.zx-tech.net/pengfeng/facebook/model"
)

var client *http.Client
var proxyUrl = `` //代理URL必须
var fbHost = `https://graph.facebook.com/`

var SaveActionLogAns = true
var SaveActionDir = `/runtime`
var SaveActionLog = true

func init() {
	proxy, _ := url.Parse(proxyUrl)
	netTransport := &http.Transport{
		DisableKeepAlives: false,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 3 * time.Minute,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       10 * time.Minute,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		Proxy:                 http.ProxyURL(proxy),
	}
	client = &http.Client{
		Transport: netTransport,
	}
	fmt.Println("初始化了！")
}

func Action(method, urls, action string, postData map[string]string, headers map[string]string) (by []byte, err error) {
	if !(strings.Contains(urls, "https:") || strings.Contains(urls, "http:")) {
		urls = fbHost + urls
	}
	var req *http.Request
	if strings.Contains(action, "json") {
		buf, _ := json.Marshal(postData)
		req, err = http.NewRequest(method, urls, bytes.NewBuffer(buf))
	} else {

		if strings.ToLower(method) == "post" {
			val := url.Values{}
			for k, v := range postData {
				val.Add(k, v)
			}
			req, err = http.NewRequest(method, urls, strings.NewReader(val.Encode()))
		} else {
			if postData != nil {
				val := url.Values{}
				for k, v := range postData {
					val.Add(k, v)
				}
				req, err = http.NewRequest(method, urls+"?"+val.Encode(), strings.NewReader(val.Encode()))
			} else {
				req, err = http.NewRequest(method, urls, nil)
			}

		}
	}
	if err != nil {
		err = fmt.Errorf("http.NewRequest is fail: %v", err.Error())
		return
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		err = fmt.Errorf("client.Do is fail: %v", err.Error())
		return
	}
	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("ioutil.ReadAll is fail: %v", err.Error())
		return
	}
	err = handErr(by)
	if SaveActionLog || err != nil {
		_, err1 := saveFile(urls, method, postData, by, err)
		if err1 != nil {
			fmt.Println(err1.Error(), "err1.Error")
		}
	}
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("post amazon HTTP CODE:" + fmt.Sprint(resp.StatusCode))
		return
	}
	return
}

func ActionAll(method, urls, action string, postData map[string]string, headers map[string]string) (by []byte, err error) {
	by, err = Action(method, urls, action, postData, headers)
	if err != nil {
		return
	}
	var res0 model.ListAll
	err = json.Unmarshal(by, &res0)
	list := res0.Data
	if err != nil {
		err = fmt.Errorf("JSON转化失败：%v", err.Error())
		return
	}
	next := res0.Paging.Next
	if next != "" {
		for {
			if next == "" {
				break
			}
			var res1 model.ListAll
			data1, err1 := Action("GET", next, "", nil, nil)
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
	res0.Data = list
	res0.Paging.Next = next
	by, err = json.Marshal(res0)
	return
}
func handErr(by []byte) (err error) {
	type er struct {
		Message   string `json:"message"`
		Type      string `json:"type"`
		Code      int    `json:"code"`
		FbtraceId string `json:"fbtrace_id"`
	}
	type res struct {
		Error er `json:"error"`
	}
	if strings.Contains(string(by), "error") {
		var obj res
		err = json.Unmarshal(by, &obj)
		if err != nil {
			return
		}
		err = fmt.Errorf("response fail: %v", obj.Error.Message)
	}
	return
}
func GetDomain(s string) string {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Get(s)
	if err == nil && (res.StatusCode == 301 || res.StatusCode == 302) {
		s = res.Header.Get("Location")
	}
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println("处理失败")
		return s
	}
	host := u.Hostname()
	ss := strings.Split(host, ".")
	if len(ss) > 2 {
		host = strings.Join(ss[1:], ".")
	}
	return host
}
