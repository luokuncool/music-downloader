package ntes

type SearchResponse struct {
	Result Result `json:"result"`
	Code   int    `json:"code"`
}
type Artists struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	PicURL    interface{}   `json:"picUrl"`
	Alias     []interface{} `json:"alias"`
	AlbumSize int           `json:"albumSize"`
	PicID     int           `json:"picId"`
	Img1V1URL string        `json:"img1v1Url"`
	Img1V1    int           `json:"img1v1"`
	Trans     interface{}   `json:"trans"`
}
type Artist struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	PicURL    interface{}   `json:"picUrl"`
	Alias     []interface{} `json:"alias"`
	AlbumSize int           `json:"albumSize"`
	PicID     int           `json:"picId"`
	Img1V1URL string        `json:"img1v1Url"`
	Img1V1    int           `json:"img1v1"`
	Trans     interface{}   `json:"trans"`
}
type Album struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Artist      Artist `json:"artist"`
	PublishTime int64  `json:"publishTime"`
	Size        int    `json:"size"`
	CopyrightID int    `json:"copyrightId"`
	Status      int    `json:"status"`
	PicID       int64  `json:"picId"`
	Mark        int    `json:"mark"`
}
type Songs struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Artists     []Artists     `json:"artists"`
	Album       Album         `json:"album"`
	Duration    int           `json:"duration"`
	CopyrightID int           `json:"copyrightId"`
	Status      int           `json:"status"`
	Alias       []interface{} `json:"alias"`
	Rtype       int           `json:"rtype"`
	Ftype       int           `json:"ftype"`
	Mvid        int           `json:"mvid"`
	Fee         int           `json:"fee"`
	RURL        interface{}   `json:"rUrl"`
	Mark        int           `json:"mark"`
}
type Result struct {
	Songs     []Songs `json:"songs"`
	SongCount int     `json:"songCount"`
}
