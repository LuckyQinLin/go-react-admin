package utils

import (
	"bytes"
	"github.com/goccy/go-json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

const IpUrl = "http://whois.pconline.com.cn/ipJson.jsp"

func IpAddress(ip string) (res string) {
	var (
		jsonObj struct {
			Ip   string `json:"ip"`
			Addr string `json:"addr"`
		}
		resp *http.Response
		body []byte
		err  error
	)
	// 判断是否为内网地址
	if ip == "::1" {
		res = "::1"
		return
	}

	if ip[:3] == "10." || ip[:4] == "172." || ip[:4] == "192." {
		res = "内网IP"
		return
	}
	if resp, err = http.Get(IpUrl + "?ip=" + ip + "&json=true"); err != nil {
		res = "未知"
		return
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		res = "未知"
		return
	}
	if body, err = io.ReadAll(transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())); err != nil {
		res = "未知"
		return
	}
	if err = json.Unmarshal(body, &jsonObj); err != nil {
		res = "未知"
		return
	}
	return jsonObj.Addr
}
