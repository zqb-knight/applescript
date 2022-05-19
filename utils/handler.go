package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const mapUrl = `<a href="https://restapi.amap.com/v3/staticmap?location=%s&zoom=10&size=1920*1080&markers=mid,,A:%s&key=%s">链接文本</a>`
const locationUrl = `https://restapi.amap.com/v3/geocode/regeo`
const mkey = `11a5f2853f8f05a7bf5987a57b590cd0`

var errCode = map[int]string{
	-1: "参数缺失",
	-2: "请求下游API错误",
	-3: "系统内部错误",
	1:  "成功",
}

func BuildResponse(status int, detail string) map[string]interface{} {
	return map[string]interface{}{
		"status":  status,
		"message": errCode[status],
		"detail":  detail,
	}
}

func GetDetail(lat string, lng string) map[string]interface{} {
	key := "11a5f2853f8f05a7bf5987a57b590cd0"
	myUrl := "https://restapi.amap.com/v3/geocode/regeo"
	location := lng + "," + lat
	radius := "1000"
	params := url.Values{}
	Url, err := url.Parse(myUrl)
	if err != nil {
		return BuildResponse(-2, err.Error())
	}
	params.Set("key", key)
	params.Set("location", location)
	params.Set("radius", radius)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath) // ignore_security_alert
	if err != nil {
		return BuildResponse(-2, err.Error())
	}
	if resp != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		now := time.Now().String()
		fmt.Println(now)
		fmt.Println(string(body))
		BaseSend(location, string(body))
		//log.Println(string(body))
		return BuildResponse(0, "")
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	return nil

}
