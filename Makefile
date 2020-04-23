default:
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ./build/music_downloader_windows_386.exe ./cmd/downloader
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ./build/music_downloader_windows_amd64.exe ./cmd/downloader
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./build/music_downloader_darwin_amd64 ./cmd/downloader
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ./build/music_downloader_arm ./cmd/downloader
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o ./build/music_downloader_arm64 ./cmd/downloader
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/music_downloader_amd64 ./cmd/downloader
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o ./build/music_downloader_386 ./cmd/downloader
pkg: default
	(cd ./build && tar cvf ./music_downloader.tar.gz ./*)
clean:
	find ./build -name *downloader* -exec rm {} +