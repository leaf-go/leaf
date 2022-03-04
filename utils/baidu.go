package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	BaiduAppID     = "23524951"
	BaiduAppKey    = "PRvfBsc8nZXDq7orvuwWGwuw"
	BaidusecretKey = "01DYenAbE5i1CU29REDGynFQeBAR8yhj"
)

const (
	BaiduUrlAccessToken = "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s"
	BaiduUrlCheckText   = "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined"
	BaiduUrlCheckImage  = "https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined"
)

type Baidu struct {
}

func (this *Baidu) getAccessToken(ctx *gin.Context) (accessToken string) {
	//cache := client.GetRedis()
	//
	//accessToken, err := cache.IdGet("baidu_accessToken").Result()

	if accessToken != "" {
		return
	}

	url := fmt.Sprintf(BaiduUrlAccessToken, BaiduAppKey, BaidusecretKey)

	resp, err := http.Get(url)

	if err != nil {
		//log.Info(ctx, "baidu getAccessToken request error", log.H{
		//	"err": err.Error(),
		//})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	//log.Info(ctx, "baidu getAccessToken DefaultOutput", log.H{
	//	"body": string(body),
	//})

	res := struct {
		AccessToken string `json:"access_token"`
	}{}

	if err = Json.Unmarshal(body, &res); err != nil {
		//log.Info(ctx, "baidu getAccessToken json decode error", log.H{
		//	"err": err.Error(),
		//})
		return
	}

	//err = cache.Set("baidu_accessToken", res.AccessToken, time.Hour*24*20).Err()

	//log.Info(ctx, "baidu getAccessToken set redis", log.H{
	//	"err": err,
	//})

	accessToken = res.AccessToken

	return
}

// 1：合规，2：不合规，3：疑似，4：审核失败
func (this *Baidu) CheckText(ctx *gin.Context, text string) (result int) {
	accessToken := this.getAccessToken(ctx)
	s := fmt.Sprintf("access_token=%s&text=%s", accessToken, url.QueryEscape(text))

	//log.Info(ctx, "baidu CheckText input", log.H{
	//	"s":   s,
	//	"url": BaiduUrlCheckText,
	//})

	resp, err := http.Post(BaiduUrlCheckText,
		"application/x-www-form-urlencoded", strings.NewReader(s))

	//log.Info(ctx, "baidu CheckText DefaultOutput", log.H{
	//	"err": err,
	//})

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	res := struct {
		ErrorCode      int `json:"error_code"`
		ConclusionType int `json:"conclusionType"`
	}{}

	if err = Json.Unmarshal(body, &res); err != nil {
		//log.Error(ctx, "baidu CheckText DefaultOutput json decode error", log.H{
		//	"err": err,
		//})
		return
	}

	if res.ErrorCode != 0 {
		return
	}

	result = res.ConclusionType

	return
}

// 1：合规，2：不合规，3：疑似，4：审核失败
func (this *Baidu) CheckImage(ctx *gin.Context, imgBase64 string) (result int) {
	accessToken := this.getAccessToken(ctx)
	u := fmt.Sprintf("%s?access_token=%s", BaiduUrlCheckImage, accessToken)

	//log.Info(ctx, "baidu CheckImage input", log.H{
	//	"url": u,
	//})

	s := "image=" + url.QueryEscape(imgBase64)
	resp, err := http.Post(u, "application/x-www-form-urlencoded", strings.NewReader(s))

	//log.Info(ctx, "baidu CheckImage DefaultOutput", log.H{
	//	"err": err,
	//})

	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	res := struct {
		ErrorCode      int `json:"error_code"`
		ConclusionType int `json:"conclusionType"`
	}{}

	//log.Info(ctx, "baidu CheckText DefaultOutput body", log.H{
	//	"err":  err,
	//	"body": string(body),
	//})

	if err = Json.Unmarshal(body, &res); err != nil {
		return
	}

	if res.ErrorCode != 0 {
		return
	}

	//log.Info(ctx, "baidu CheckImage DefaultOutput body", log.H{
	//	"body": string(body),
	//})

	result = res.ConclusionType

	return
}
