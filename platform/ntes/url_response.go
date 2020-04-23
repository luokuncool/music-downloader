package ntes

type UrlResponse struct {
	Data []Data `json:"data"`
	Code int    `json:"code"`
}
type Data struct {
	ID            int         `json:"id"`
	URL           string      `json:"url"`
	Br            int         `json:"br"`
	Size          int         `json:"size"`
	Md5           string      `json:"md5"`
	Code          int         `json:"code"`
	Expi          int         `json:"expi"`
	Type          string      `json:"type"`
	Gain          float64     `json:"gain"`
	Fee           int         `json:"fee"`
	Uf            interface{} `json:"uf"`
	Payed         int         `json:"payed"`
	Flag          int         `json:"flag"`
	CanExtend     bool        `json:"canExtend"`
	FreeTrialInfo interface{} `json:"freeTrialInfo"`
	Level         string      `json:"level"`
	EncodeType    string      `json:"encodeType"`
}
