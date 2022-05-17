package utils

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

var errCode = map[int]string{
	-1: "参数缺失",
	-2: "请求下游API错误",
	-3: "系统内部错误",
	1:  "成功",
}

func BuildResponse(c *gin.Context, status int, detail string) {
	if status == -1 {
		c.JSON(200, gin.H{
			"status":  status,
			"message": errCode[status],
			"detail":  detail,
		})
	}
}

func GetDetail(c *gin.Context, lat string, lng string) map[string]interface{} {
	key := "11a5f2853f8f05a7bf5987a57b590cd0"
	myUrl := "https://restapi.amap.com/v3/geocode/regeo"
	location := lng + "," + lat
	radius := "100"
	params := url.Values{}
	Url, err := url.Parse(myUrl)
	if err != nil {
		BuildResponse(c, -2, err.Error())
		return nil
	}
	params.Set("key", key)
	params.Set("location", location)
	params.Set("radius", radius)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath) // ignore_security_alert
	if err != nil {
		BuildResponse(c, -2, err.Error())
	}
	if resp != nil {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
		BuildResponse(c, 0, "")
	}
	if resp != nil {
		err := resp.Body.Close()
		if err != nil {

		}
	}
	return nil

}
