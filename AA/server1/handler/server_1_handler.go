package handler

import (
	"AA/dto"
	"AA/server1/helpers"
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type RequestPayload struct {
	Url            string
	AnotherPayLoad string
}

func HanderSlashServer1Route(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var inputDataReceived RequestPayload

	err := decoder.Decode(&inputDataReceived)
	if err != nil {
		w.Write([]byte("No Data Received,"))
	}
	requiredUrl := inputDataReceived.Url
	var producturl string = string(requiredUrl)

	htmlresp := helpers.Crawlit(producturl)
	title := helpers.Extractit(htmlresp, "<title>(.*)<\\/title>", "<[^>]*>", 0)
	description := helpers.Extractit(htmlresp, "<p>(.*?)\n", "<[^>]*>", 2)
	reviews := helpers.Extractit(htmlresp, "<span id=\"acrCustomerReviewText\" class=\"a-size-base\">(.*?)<\\/span>", "<[^>]*>", 0)
	thumbnail := helpers.Extractit(htmlresp, "\"thumb\":\"(.*?)\"", "\"thumb\":", 0)

	crawlInfo := dto.CrawlInfo{
		URL:       requiredUrl,
		Crawltime: time.Now().Format("2006-01-02 15:04:05"),
		ProductInfo: dto.Product{
			Title:       title,
			Thumbnail:   thumbnail,
			Reviews:     reviews,
			Description: description,
		},
	}

	crawlInfoBuffer := new(bytes.Buffer)
	json.NewEncoder(crawlInfoBuffer).Encode(crawlInfo)

	req, err := http.NewRequest("POST", "http://server2:8002/server2", crawlInfoBuffer)
	if err != nil {
		logrus.Error(err)
		w.Write([]byte("Failed to create request object"))
		return
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		w.Write([]byte("Call to server 2 failed"))
		return
	}
	defer res.Body.Close()

	crawljs, err := json.Marshal(crawlInfo)
	if err != nil {
		logrus.Error(err)
	}
	w.Write(crawljs)
	return

}
