package platform

//MusicPlatform 音乐平台接口
type MusicPlatform interface {
	Search(keyword string, page string) []Song
	Url(id string) Url
}
