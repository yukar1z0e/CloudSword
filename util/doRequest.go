package util

import (
	"CloudSword/structure"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

func DoGet(url string, time time.Duration) (respBody structure.ResponseBody, err error) {
	tr := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time)
			if err != nil {
				//fmt.Println("dail timeout", err)
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	response, err := client.Do(req)
	if err != nil {
		respBody.RespCode = 404
		respBody.RespContent = ""
		fmt.Println(err)
		return respBody, err
	}
	defer response.Body.Close()
	tmpBody := response.Body
	content, _ := ioutil.ReadAll(tmpBody)
	respBody.RespCode = response.StatusCode
	respBody.RespContent = string(content)
	return respBody, err
}

func DoPost(url string, postContent string, time time.Duration) (respBody structure.ResponseBody, err error) {
	tr := &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time)
			if err != nil {
				//fmt.Println("dail timeout", err)
				return nil, err
			}
			return c, nil
		},
		MaxIdleConnsPerHost:   10,
		ResponseHeaderTimeout: time,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest("POST", url, strings.NewReader(postContent))
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:76.0) Gecko/20100101 Firefox/76.0")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	response, err := client.Do(req)
	if err != nil {
		respBody.RespCode = 404
		respBody.RespContent = ""
		return respBody, err
	}
	defer response.Body.Close()
	tmpBody := response.Body
	content, _ := ioutil.ReadAll(tmpBody)
	respBody.RespCode = response.StatusCode
	respBody.RespContent = string(content)
	return respBody, err
}
