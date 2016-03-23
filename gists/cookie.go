package main

import (
	"fmt"
	"net/http"
	"crypto/md5"
	"io/ioutil"
//	"bytes"
	"net/http/cookiejar"
	"log"
"net/url"
"strings"
)

//https://github.com/mawenbao/music-api-server/blob/master/netease_service.go
const(
	gNeteaseAPIUrlBase        = "http://music.163.com/api"
)


var (

	gNeteaseClient      = &http.Client{}
)

func unused( ... interface{}){}
func init() {
	// init netease http client
	cookies, err := cookiejar.New(nil)
	if nil != err {
		log.Fatal("failed to init netease httpclient cookiejar: %s", err)
	}
	apiUrl, err := url.Parse(gNeteaseAPIUrlBase)
	if nil != err {
		log.Fatal("failed to parse netease api url %s: %s", gNeteaseAPIUrlBase, err)
	}

	unused(cookies, apiUrl)
	// netease api requires some cookies to work
	cookies.SetCookies(apiUrl, []*http.Cookie{
		&http.Cookie{Name: "appver", Value: "1.4.1.62460"},
		&http.Cookie{Name: "os", Value: "pc"},
		&http.Cookie{Name: "osver", Value: "Microsoft-Windows-7-Ultimate-Edition-build-7600-64bit"},
	})
	gNeteaseClient.Jar = cookies
}

//
//func GetUrl(client *http.Client, url string) []byte {
//	cacheKey := GenUrlCacheKey(url)
//	if "" == cacheKey {
//		return nil
//	}
//	// try to load from cache first
//	body := GetCache(cacheKey)
//	if nil != body {
//		return body
//	}
//
//	// cache missed, do http request
//	resp, err := client.Get(url)
//	if err != nil {
//		log.Printf("error get url %s: %s", url, err)
//		return nil
//	}
//	defer resp.Body.Close()
//	body, err = ioutil.ReadAll(resp.Body)
//	if err != nil {
//		log.Printf("error getting response body from url %s: %s", url, err)
//		return nil
//	}
//	// update cache, with compression
//	SetCache(cacheKey, body, 60*60*60)
//	return body
//}



func main() {
	fmt.Println("hello netease music")


	m := md5.New()
	m.Write([]byte("LEWGUN"))

	body := map[string]string {
		"password": string(""),
		"rememberLogin": "true",
		"username": "",
	}

	encBody, _ := encryptRequest(body)

	//req, _ := http.NewRequest("GET", "http://music.163.com/api/radio/get", nil )
	req, _ := http.NewRequest("POST",
		"https://music.163.com/weapi/login/cellphone",
		strings.NewReader(fmt.Sprintf("params=%s&encSecKey=%s",encBody["params"], encBody["encSecKey"])))

	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip,deflate,sdch")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4*")
	req.Header.Add("Connection", "application/x-www-form-urlencoded")
	req.Header.Add("Host", "music.163.com")
	req.Header.Add("Referer", "http://music.163.com/search/")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36")

	resp, _ := gNeteaseClient.Do(req)

	defer resp.Body.Close()
	d, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("get value: ")
	fmt.Println(string(d), len(d))

	fmt.Println(resp.Header, resp.StatusCode)
	fmt.Println(resp.Cookies())


	////
	////r, _ := http.Get("http://music.163.com/api/radio/get")
	//defer resp.Body.Close()
	//
	//d, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(d))

}
