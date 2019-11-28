package kugou

import "music-downloader/platform"

//Kugou 酷狗音乐支持
type Kugou struct {
}

func (k *Kugou) Search(keyword string, page string) []platform.Song {
	return nil
}

func (k *Kugou) Url(id string) platform.Url {
	return platform.Url{
		BitRate: "",
		Url:     "",
		Size:    0,
	}
}
