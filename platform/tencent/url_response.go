package tencent

type UrlResponse struct {
	Req0 struct {
		Data struct {
			Expiration int    `json:"expiration"`
			LoginKey   string `json:"login_key"`
			Midurlinfo []struct {
				CommonDownfromtag int    `json:"common_downfromtag"`
				Errtype           string `json:"errtype"`
				Filename          string `json:"filename"`
				Flowfromtag       string `json:"flowfromtag"`
				Flowurl           string `json:"flowurl"`
				Hisbuy            int    `json:"hisbuy"`
				Hisdown           int    `json:"hisdown"`
				Isbuy             int    `json:"isbuy"`
				Isonly            int    `json:"isonly"`
				Onecan            int    `json:"onecan"`
				Opi128Kurl        string `json:"opi128kurl"`
				Opi192Koggurl     string `json:"opi192koggurl"`
				Opi192Kurl        string `json:"opi192kurl"`
				Opi30Surl         string `json:"opi30surl"`
				Opi48Kurl         string `json:"opi48kurl"`
				Opi96Kurl         string `json:"opi96kurl"`
				Opiflackurl       string `json:"opiflackurl"`
				P2Pfromtag        int    `json:"p2pfromtag"`
				Pdl               int    `json:"pdl"`
				Pneed             int    `json:"pneed"`
				Pneedbuy          int    `json:"pneedbuy"`
				Premain           int    `json:"premain"`
				Purl              string `json:"purl"`
				Qmdlfromtag       int    `json:"qmdlfromtag"`
				Result            int    `json:"result"`
				Songmid           string `json:"songmid"`
				Tips              string `json:"tips"`
				UIAlert           int    `json:"uiAlert"`
				VipDownfromtag    int    `json:"vip_downfromtag"`
				Vkey              string `json:"vkey"`
				Wififromtag       string `json:"wififromtag"`
				Wifiurl           string `json:"wifiurl"`
			} `json:"midurlinfo"`
			Msg          string   `json:"msg"`
			Retcode      int      `json:"retcode"`
			Servercheck  string   `json:"servercheck"`
			Sip          []string `json:"sip"`
			Testfile2G   string   `json:"testfile2g"`
			Testfilewifi string   `json:"testfilewifi"`
			Thirdip      []string `json:"thirdip"`
			Uin          string   `json:"uin"`
			VerifyType   int      `json:"verify_type"`
		} `json:"data"`
		Code int `json:"code"`
	} `json:"req_0"`
	Code int   `json:"code"`
	Ts   int64 `json:"ts"`
}
