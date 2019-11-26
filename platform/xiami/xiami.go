package xiami

import "music-downloader/platform"

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
