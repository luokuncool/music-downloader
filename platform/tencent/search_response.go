package tencent

type SearchResponse struct {
	Code int `json:"code"`
	Data struct {
		Keyword  string        `json:"keyword"`
		Priority int           `json:"priority"`
		Qc       []interface{} `json:"qc"`
		Semantic struct {
			Curnum   int           `json:"curnum"`
			Curpage  int           `json:"curpage"`
			List     []interface{} `json:"list"`
			Totalnum int           `json:"totalnum"`
		} `json:"semantic"`
		Song struct {
			Curnum  int `json:"curnum"`
			Curpage int `json:"curpage"`
			List    []struct {
				Action struct {
					Alert  int `json:"alert"`
					Icons  int `json:"icons"`
					Msg    int `json:"msg"`
					Switch int `json:"switch"`
				} `json:"action"`
				Album struct {
					ID           int    `json:"id"`
					Mid          string `json:"mid"`
					Name         string `json:"name"`
					Pmid         string `json:"pmid"`
					Subtitle     string `json:"subtitle"`
					Title        string `json:"title"`
					TitleHilight string `json:"title_hilight"`
				} `json:"album"`
				Chinesesinger int    `json:"chinesesinger"`
				Desc          string `json:"desc"`
				DescHilight   string `json:"desc_hilight"`
				Docid         string `json:"docid"`
				File          struct {
					B30S        int    `json:"b_30s"`
					E30S        int    `json:"e_30s"`
					MediaMid    string `json:"media_mid"`
					Size128     int    `json:"size_128"`
					Size128Mp3  int    `json:"size_128mp3"`
					Size320     int    `json:"size_320"`
					Size320Mp3  int    `json:"size_320mp3"`
					SizeAac     int    `json:"size_aac"`
					SizeApe     int    `json:"size_ape"`
					SizeDts     int    `json:"size_dts"`
					SizeFlac    int    `json:"size_flac"`
					SizeOgg     int    `json:"size_ogg"`
					SizeTry     int    `json:"size_try"`
					StrMediaMid string `json:"strMediaMid"`
					TryBegin    int    `json:"try_begin"`
					TryEnd      int    `json:"try_end"`
				} `json:"file"`
				Fnote      int           `json:"fnote"`
				Genre      int           `json:"genre"`
				Grp        []interface{} `json:"grp"`
				ID         int           `json:"id"`
				IndexAlbum int           `json:"index_album"`
				IndexCd    int           `json:"index_cd"`
				Interval   int           `json:"interval"`
				Isonly     int           `json:"isonly"`
				Ksong      struct {
					ID  int    `json:"id"`
					Mid string `json:"mid"`
				} `json:"ksong"`
				Language     int    `json:"language"`
				Lyric        string `json:"lyric"`
				LyricHilight string `json:"lyric_hilight"`
				Mid          string `json:"mid"`
				Mv           struct {
					ID  int    `json:"id"`
					Vid string `json:"vid"`
				} `json:"mv"`
				Name      string `json:"name"`
				NewStatus int    `json:"newStatus"`
				Nt        int64  `json:"nt"`
				Pay       struct {
					PayDown    int `json:"pay_down"`
					PayMonth   int `json:"pay_month"`
					PayPlay    int `json:"pay_play"`
					PayStatus  int `json:"pay_status"`
					PriceAlbum int `json:"price_album"`
					PriceTrack int `json:"price_track"`
					TimeFree   int `json:"time_free"`
				} `json:"pay"`
				Pure   int `json:"pure"`
				Singer []struct {
					ID           int    `json:"id"`
					Mid          string `json:"mid"`
					Name         string `json:"name"`
					Title        string `json:"title"`
					TitleHilight string `json:"title_hilight"`
				} `json:"singer"`
				Status       int    `json:"status"`
				Subtitle     string `json:"subtitle"`
				T            int    `json:"t"`
				Tag          int    `json:"tag"`
				TimePublic   string `json:"time_public"`
				Title        string `json:"title"`
				TitleHilight string `json:"title_hilight"`
				Type         int    `json:"type"`
				URL          string `json:"url"`
				Ver          int    `json:"ver"`
				Volume       struct {
					Gain float64 `json:"gain"`
					Lra  float64 `json:"lra"`
					Peak float64 `json:"peak"`
				} `json:"volume"`
				Format string `json:"format,omitempty"`
			} `json:"list"`
			Totalnum int `json:"totalnum"`
		} `json:"song"`
		Tab       int           `json:"tab"`
		Taglist   []interface{} `json:"taglist"`
		Totaltime int           `json:"totaltime"`
		Zhida     struct {
			Chinesesinger int `json:"chinesesinger"`
			Type          int `json:"type"`
		} `json:"zhida"`
	} `json:"data"`
	Message string `json:"message"`
	Notice  string `json:"notice"`
	Subcode int    `json:"subcode"`
	Time    int    `json:"time"`
	Tips    string `json:"tips"`
}
