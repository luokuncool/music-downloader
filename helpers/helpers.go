package helpers

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var httpClient http.Client

func HttpGet(url string, body url.Values, header http.Header) string {
	rq, _ := http.NewRequest("GET", url, strings.NewReader(""))
	if header != nil {
		rq.Header = header
	}
	rq.URL.RawQuery = body.Encode()

	resp, e := httpClient.Do(rq)
	ErrorCheck(e)

	out, e := ioutil.ReadAll(resp.Body)
	ErrorCheck(e)

	return string(out)
}

func HttpPost(url string, body url.Values, header http.Header) string {
	rq, e := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
	ErrorCheck(e)

	if header != nil {
		rq.Header = header
	}

	resp, e := httpClient.Do(rq)
	ErrorCheck(e)

	out, e := ioutil.ReadAll(resp.Body)
	ErrorCheck(e)

	return string(out)
}

func ErrorCheck(e error) {
	if e != nil {
		panic(errors.WithStack(e))
	}
}
