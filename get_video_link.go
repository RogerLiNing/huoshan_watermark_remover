package huoshan

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type JsonData struct {
	Data Data `json:"data"`
}

type Data struct {
	ItemInfo ItemInfo `json:"item_info"`
}

type ItemInfo struct {
	URL string `json:"url"`
}

func getVideoLink(id string) (string, error) {

	client := &http.Client{}
	// 通过这个接口获取视频信息，其中包括带有水印的链接
	url := "https://share.huoshan.com/api/item/info?item_id=" + id

	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "*/*")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Connection", "keep-alive")

	response, err := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonByteData := []byte(string(body))

	jsonData := JsonData{}
	json.Unmarshal(jsonByteData, &jsonData)
	var videoLink = ""

	if len(jsonData.Data.ItemInfo.URL) > 0 {
		videoLink = strings.Replace(jsonData.Data.ItemInfo.URL, "reflow", "source", 1)
		videoLink = strings.Replace(videoLink, "watermark=2", "watermark=0", 1)
	}

	return videoLink, nil

}
