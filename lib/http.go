package lib

import (
	"io/ioutil"
	"net/http"
)

func GetHTML(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	html := string(body)

	return html, nil
}
