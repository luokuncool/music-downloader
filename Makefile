default:
	env CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./build/music_downloader_windows_386.exe
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/music_downloader_windows_amd64.exe
	env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/music_downloader_darwin_amd64
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/music_downloader_downloader_arm
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./build/music_downloader_arm64