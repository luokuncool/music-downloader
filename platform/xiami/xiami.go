package xiami

import (
	"music-downloader/platform"
)

//Xiami 虾米音乐支持
type Xiami struct {
}

func (x *Xiami) Search(keyword string, page string) []platform.Song {
	return nil
}

func (x *Xiami) Url(id string) platform.Url {
	return platform.Url{
		BitRate: "",
		Url:     "",
		Size:    0,
	}
}
