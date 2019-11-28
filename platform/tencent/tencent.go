package tencent

import (
	"encoding/json"
	"github.com/fatih/structs"
	"math/rand"
	"music-downloader/helpers"
	"music-downloader/platform"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Tencent struct {
}

var header = http.Header{
	"Referer":         {"http://y.qq.com"},
	"Cookie":          {"pgv_pvi=22038528; pgv_si=s3156287488; pgv_pvid=5535248600; yplayer_open=1; ts_last=y.qq.com/portal/player.html; ts_uid=4847550686; yq_index=0; qqmusic_fromtag=66; player_exist=1"},
	"User-Agent":      {"QQ%E9%9F%B3%E4%B9%90/54409 CFNetwork/901.1 Darwin/17.6.0 (x86_64)"},
	"Accept":          {"*/*"},
	"Accept-Language": {"zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4"},
	"Connection":      {"keep-alive"},
	"Content-Type":    {"application/x-www-form-urlencoded"},
}

func (t *Tencent) Search(keyword string, page string) []platform.Song {
	if page == "" {
		page = "1"
	}
	q := url.Values{
		"format":   {"json"},
		"p":        {page},
		"n":        {"10"},
		"w":        {keyword},
		"aggr":     {"1"},
		"lossless": {"1"},
		"cr":       {"1"},
		"new_json": {"1"},
	}
	out := helpers.HttpGet("https://c.y.qq.com/soso/fcgi-bin/client_search_cp", q, header)
	responseObj := SearchResponse{}
	json.Unmarshal([]byte(out), &responseObj)

	var result []platform.Song
	if len(responseObj.Data.Song.List) == 0 {
		return result
	}
	for _, value := range responseObj.Data.Song.List {
		var singers []string
		for _, singer := range value.Singer {
			singers = append(singers, singer.Name)
		}
		song := platform.Song{
			Singer:    strings.Join(singers, ","),
			Title:     value.Title,
			Url:       t.Url(value.Mid),
			AlbumName: value.Album.Name,
		}
		if song.Url.Url == "" || song.Title == "" {
			continue
		}
		result = append(result, song)
		//fmt.Println(song.Title, t.Url(value.Mid))
	}

	//f, _ := os.Create("search-result.json")
	//f.Write(out)
	//f.Close()

	return result
}

// Url 获取歌曲的下载链接
func (t *Tencent) Url(id string) platform.Url {
	songQuery := url.Values{"songmid": {id}, "platform": {"yqq"}, "format": {"json"}}
	out := helpers.HttpGet("https://c.y.qq.com/v8/fcg-bin/fcg_play_single_song.fcg", songQuery, header)
	result := SongResponse{}
	json.Unmarshal([]byte(out), &result)

	guid := strconv.FormatUint(rand.Uint64()%10000000000, 10)
	payload := map[string]interface{}{
		"req_0": map[string]interface{}{
			"module": "vkey.GetVkeyServer",
			"method": "CgiGetVkey",
			"param": map[string]interface{}{
				"guid":      guid,
				"songmid":   []string{},
				"filename":  []string{},
				"songtype":  []string{},
				"uin":       "0",
				"loginflag": 1,
				"platform":  "20",
			},
		},
	}

	var songmid []string
	var filename []string
	var songtype []int
	types := [][]string{
		{"Size320Mp3", "320", "M800", "mp3"},
		{"Size192Aac", "192", "C600", "m4a"},
		{"Size128Mp3", "128", "M500", "mp3"},
		{"Size96Aac", "96", "C400", "m4a"},
		{"Size48Aac", "48", "C200", "m4a"},
		{"Size24Aac", "24", "C100", "m4a"},
	}
	for _, value := range types {
		songmid = append(songmid, result.Data[0].Mid)
		filename = append(filename, value[2]+result.Data[0].File.MediaMid+"."+value[3])
		songtype = append(songtype, result.Data[0].Type)
	}

	param := payload["req_0"].(map[string]interface{})["param"].(map[string]interface{})
	param["songtype"] = songtype
	param["filename"] = filename
	param["songmid"] = songmid

	payloadJson, _ := json.Marshal(payload)

	q := url.Values{}
	q.Add("needNewCode", "0")
	q.Add("platform", "yqq.json")
	q.Add("format", "json")
	q.Add("data", string(payloadJson))

	urlResp := UrlResponse{}
	urlOut := helpers.HttpGet("https://u.y.qq.com/cgi-bin/musicu.fcg", q, header)
	json.Unmarshal([]byte(urlOut), &urlResp)

	//f, _ := os.Create("url-result.json")
	//f.Write(urlOut)
	//f.Close()

	vkeys := urlResp.Req0.Data.Midurlinfo
	for index, vo := range types {
		file := structs.Map(result.Data[0].File)
		if vkeys[index].Vkey != "" && file[vo[0]].(int) > 0 {
			return platform.Url{
				Url:     urlResp.Req0.Data.Sip[0] + vkeys[index].Purl,
				Size:    file[vo[0]].(int),
				BitRate: vo[1],
			}
		}
	}

	return platform.Url{
		BitRate: "",
		Url:     "",
		Size:    0,
	}
}
