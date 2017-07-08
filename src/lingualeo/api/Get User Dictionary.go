package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func sendGetUserDictionary() {
	// Get User Dictionary (POST https://lingualeo.com/userdict/json)

	params := url.Values{}
	params.Set("groupId", "dictionary")
	params.Set("sortBy", "date")
	params.Set("filter", "all")
	params.Set("wordType", "0")
	params.Set("page", "12")
	body := bytes.NewBufferString(params.Encode())

	// Create client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", "https://lingualeo.com/userdict/json", body)

	// Headers
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "https://lingualeo.com")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")
	req.Header.Add("Cookie", "optimizelyEndUserId=oeu1480257581038r0.9517318914113231; _ym_uid=1480257582117317868; browser-plugins-msg-hide=1; lingualeouid=1493065545674622; userid=11981; mobilePanelChecked=1; v_cnt=1; daily_notification=1; _ym_isad=2; servid=1f01000a944960bf7c451ee21724f83dbe9c916d4607a8933e02eda660c6a390361bfce9552c6a7b; remember=cd2e000088d15423578b6f3147e8d64fa4c220ae16854f485d9cc1e42264d8b30ff35f21323b4806; AWSELB=81E385C912B24FF6E77010EAF57B6254C98AA295E079117979E2AB1017DEC24203BCB4035817165AE4A22D087780E683EE96775352406ADBF019245FA4BC3FD513C71C8ACCC3B3791E4CCA8EB6888C9744E96B1EA9; optimizelySegments=%7B%221027890468%22%3A%22direct%22%2C%221036670278%22%3A%22gc%22%2C%221037430770%22%3A%22false%22%7D; optimizelyBuckets=%7B%7D; browser-plugins-msg-start-show=1; _ga=GA1.2.257556441.1480257582; __asc=1cf211c415ba7046171a496373d; __auc=8ecb26ea158a63a325a4d3fbc72; _ym_visorc_837359=w")
	req.Header.Add("Referer", "https://lingualeo.com/ru/glossary/learn/dictionary")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "ru,uk;q=0.8,en-US;q=0.6,en;q=0.4")

	// Fetch Request
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	// Read Response Body
	respBody, _ := ioutil.ReadAll(resp.Body)

	// Display Results
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response Body : ", string(respBody))
}
