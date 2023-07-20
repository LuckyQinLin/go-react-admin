package utils

import (
	"github.com/goccy/go-json"
	"io"
	"net/http"
)

const IpUrl = "http://whois.pconline.com.cn/ipJson.jsp"

func IpAddress(ip string) string {
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

	if ip[:3] == "10." || ip[:4] == "172." || ip[:4] == "192." {
		return "内网IP"
	}
	if resp, err = http.Get(IpUrl + "?ip=" + ip + "&json=true"); err != nil {
		return "未知"
	}
	defer resp.Body.Close()
	if body, err = io.ReadAll(resp.Body); err != nil {
		return "未知"
	}
	if err = json.Unmarshal(body, &jsonObj); err != nil {
		return "未知"
	}
	return jsonObj.Addr
}
