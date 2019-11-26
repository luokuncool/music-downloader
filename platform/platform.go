package platform

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type MusicPlatform interface {
	Search(keyword string, page string) []Song
	Url(id string) Url
}

func Get(url string, header http.Header, body url.Values, response interface{}) {
	rq, _ := http.NewRequest("GET", url, strings.NewReader(body.Encode()))
	rq.Header = header

	c := http.Client{}
	resp, _ := c.Do(rq)
	output, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(output, &response)
}
