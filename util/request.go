package util

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

//get res.code，res.header
func HttpRequest(url string) (response *http.Response, err error) {
	tr := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second)
			if err != nil {
				//fmt.Println("dail timeout", err)
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	response, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return response, err
}

//get res.body
func GetContent(url string) (response []byte, err error) {
	tr := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*2)
			if err != nil {
				//fmt.Println("dail timeout", err)
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time.Second * 2,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	return response, err
}

func Test(url string) {
	fmt.Println("start test httpRequest")
	res, _ := HttpRequest(url)
	fmt.Println(res.Header)
	for _, v := range res.Header {
		fmt.Println(v)
		for _, d := range v {
			if strings.Contains(d, "thinkphp") {
				fmt.Println("dd")
				break
			}
		}
	}
	//fmt.Println(res.StatusCode)

	//测试md5流计算
	/*res1 := GetContent(url)
	fmt.Println(string(res1))
	md5Byte := md5.Sum(res1)
	md5Value := fmt.Sprintf("%x", md5Byte)
	fmt.Println(md5Value)*/
	fmt.Println("start test getContent")
}
