package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Url  string
	Size int
}

func GetContent(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// 고루틴으로 실행할 함수
func GetContentLength(url string, channel chan Page) {
	body, err := GetContent(url)
	if err != nil {
		log.Fatal(err)
	}

	// 채널에 값을 전달
	channel <- Page{url, len(body)}
}
