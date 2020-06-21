package helpers

import (
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func Crawlit(url string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logrus.Error(err)
	}
	//Defining Headers for Get request
	req.Header.Set("Host", "www.amazon.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	//Making a get request to amazon which returns the response
	resp, err := client.Do(req)

	if err != nil {
		logrus.Error(err)
	}
	//Converting Bytestream Response Content to String
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Info("Error")
	}
	htmlresp := string(body)
	return htmlresp
}
