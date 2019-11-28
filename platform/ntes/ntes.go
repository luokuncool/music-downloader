package ntes

import "music-downloader/platform"

//Ntes 网易云音乐支持
type Ntes struct {
}

//Search 网易云音乐搜索
func (n *Ntes) Search(keyword string, page string) []platform.Song {
	return nil
}

//Url 获取网易云音乐歌曲下载链接
func (n *Ntes) Url(id string) platform.Url {
	return platform.Url{
		BitRate: "",
		Url:     "",
		Size:    0,
	}
}
