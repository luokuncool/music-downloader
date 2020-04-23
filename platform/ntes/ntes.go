package ntes

import (
	"encoding/json"
	"fmt"
	"github.com/WithLin/ncma/utils"
	"math/rand"
	"music-downloader/platform"
	u "music-downloader/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

//Ntes 网易云音乐支持
type Ntes struct {
}

var header = http.Header{
	"Referer":         {"https://music.163.com/"},
	"Cookie":          {"appver=1.5.9; os=osx; __remember_me=true; osver=%E7%89%88%E6%9C%AC%2010.13.5%EF%BC%88%E7%89%88%E5%8F%B7%2017F77%EF%BC%89;"},
	"User-Agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/605.1.15 (KHTML, like Gecko)"},
	"X-Real-IP":       {long2ip(randIp(1884815360, 1884890111))},
	"Accept":          {"*/*"},
	"Accept-Language": {"zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4"},
	"Connection":      {"keep-alive"},
	"Content-Type":    {"application/x-www-form-urlencoded"},
}

func long2ip(ipInt int64) string {
	// need to do two bit shifting and “0xff” masking
	b0 := strconv.FormatInt((ipInt>>24)&0xff, 10)
	b1 := strconv.FormatInt((ipInt>>16)&0xff, 10)
	b2 := strconv.FormatInt((ipInt>>8)&0xff, 10)
	b3 := strconv.FormatInt(ipInt&0xff, 10)
	return b0 + "." + b1 + "." + b2 + "." + b3
}

func randIp(min int64, max int64) int64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Int63n(max-min) + min
}

//Search 网易云音乐搜索
func (n *Ntes) Search(keyword string, page string) []platform.Song {
	pageInt, _ := strconv.Atoi(page)
	params := fmt.Sprintf(`{"s":"%s","type":1,"limit":10,"total":true,"csrf_token":"","offset":%d}`, keyword, (pageInt-1)*10)
	encText, encSecKey, _ := utils.Encrypt(params)
	res := u.HttpPost("https://music.163.com/weapi/search/get", url.Values{"params": {encText}, "encSecKey": {encSecKey}}, header)

	responseObj := SearchResponse{}
	json.Unmarshal([]byte(res), &responseObj)

	var result []platform.Song
	for _, value := range responseObj.Result.Songs {
		var singers []string
		for _, singer := range value.Artists {
			singers = append(singers, singer.Name)
		}
		song := platform.Song{
			Singer:    strings.Join(singers, ","),
			Title:     value.Name,
			Url:       n.Url(strconv.Itoa(value.ID)),
			AlbumName: value.Album.Name,
		}
		if song.Url.Url == "" || song.Title == "" {
			continue
		}
		result = append(result, song)
	}
	return result
}

//Url 获取网易云音乐歌曲下载链接
func (n *Ntes) Url(id string) platform.Url {
	params := fmt.Sprintf(`{"ids":[%s],"br":320000,"csrf_token":""}`, id)
	encText, encSecKey, _ := utils.Encrypt(params)
	res := u.HttpPost("http://music.163.com/weapi/song/enhance/player/url", url.Values{"params": {encText}, "encSecKey": {encSecKey}}, header)

	responseObj := UrlResponse{}
	json.Unmarshal([]byte(res), &responseObj)

	for _, data := range responseObj.Data {
		return platform.Url{
			Url:     data.URL,
			Size:    data.Size,
			BitRate: strconv.Itoa(data.Br),
		}
	}

	return platform.Url{
		BitRate: "",
		Url:     "",
		Size:    0,
	}
}
