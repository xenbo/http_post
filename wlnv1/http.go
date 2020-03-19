package wlnv1

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func PostData(addr string, argv string) string {

	payload := strings.NewReader(argv)
	req, _ := http.NewRequest("POST", addr, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=utf-8")
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}

func Get(addr string, argv string)  {

}