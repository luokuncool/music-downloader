package main

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/olekukonko/tablewriter"
	"github.com/peterh/liner"
	"github.com/urfave/cli"
	"log"
	"music-downloader/platform"
	"music-downloader/platform/tencent"
	"music-downloader/platform/xiami"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var (
	historyFn   = filepath.Join(os.TempDir(), ".liner_example_history")
	commands    = []string{"search", "search -platform=xiami", "download", "download all", "page", "page next", "page prev"}
	currentPage = int64(1)
	platforms   = map[string]platform.MusicPlatform{
		"tencent": &tencent.Tencent{},
		"xiami":   &xiami.Xiami{},
	}
	searchResultData []platform.Song
	keyword          string
	musicPlatform    platform.MusicPlatform
)

func main() {
	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commands {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	if f, err := os.Open(historyFn); err == nil {
		line.ReadHistory(f)
		f.Close()
	}

inputLine:
	if input, err := line.Prompt("Music-Downloader$ "); err == nil {
		input = strings.TrimSpace(input)
		if input == "" {
			goto inputLine
		}

		cli.OsExiter = func(code int) {}
		app := &cli.App{}
		app.Name = "music-downloader"
		app.Commands = []cli.Command{
			{
				Name:        "search",
				Description: "输入关键字搜索歌曲",
				Usage:       "search [关键字] [音乐平台]",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "platform",
						Value: "tencent",
					},
				},
				Action: func(c *cli.Context) {
					keyword = c.Args().Get(0)
					p := c.String("platform")
					if p == "" {
						p = "tencent"
					}
					musicPlatform = platforms[p]
					Search(c.Args().Get(0), "1")
				},
			},
			{
				Name:        "page",
				Description: "跳转到搜索结果的指定页",
				Usage:       "page [页数]",
				Action: func(c *cli.Context) {
					page := c.Args().Get(0)
					if page == "next" {
						page = strconv.FormatInt(currentPage+1, 10)
					} else if page == "prev" {
						page = strconv.FormatInt(currentPage-1, 10)
					}
					currentPage, _ = strconv.ParseInt(page, 10, 64)
					Search(keyword, page)
				},
			},
			{
				Name:        "download",
				Description: "下载搜索结果中的指定序号歌曲，多个可用英文逗号隔开",
				Usage:       "download [序号]",
				Action: func(c *cli.Context) {
					serialNumbers := strings.Split(c.Args().Get(0), ",")
					if serialNumbers[0] == "all" {
						serialNumbers = []string{}
						for key := range searchResultData {
							serialNumbers = append(serialNumbers, strconv.FormatInt(int64(key), 10))
						}
					}
					for _, serialNumber := range serialNumbers {
						idx, _ := strconv.ParseUint(serialNumber, 10, 64)
						if l := len(searchResultData); idx >= 0 && idx < uint64(l) {
							Download(searchResultData[idx])
						}
					}
				},
			},
		}
		app.Description = `这是一个音乐检索及下载的小工具`
		pattern := `("[^"]*"|[^"\s]+)(\s+|$)`
		r := regexp.MustCompile(pattern)
		args := []string{""}
		for _, arg := range r.FindAllString(input, -1) {
			args = append(args, strings.TrimSpace(arg))
		}
		err := app.Run(args)
		if err != nil {
			//log.Fatal(err)
		}
		line.AppendHistory(input)
	} else if err == liner.ErrPromptAborted {
		log.Print("Aborted")
		return
	} else {
		log.Print("Error reading line: ", err)
	}

	if f, err := os.Create(historyFn); err != nil {
		log.Print("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
	goto inputLine
}

func Search(keyword string, page string) {
	if keyword == "" {
		fmt.Println("请输入关键字检索！")
		return
	}
	searchResultData = musicPlatform.Search(keyword, page)

	//out, _ := json.Marshal(searchResultData)
	//f, _ := os.Create("search-result2.json")
	//f.Write(out)
	//f.Close()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{"序号", "歌名", "歌手", "专辑", "大小", "比特率"})
	for i, v := range searchResultData {
		row := []string{
			strconv.Itoa(i), v.Title, v.Singer, v.AlbumName,
			fmt.Sprintf("%.2fM", float64(v.Url.Size/1024/1024)),
			v.Url.BitRate,
		}
		table.Append(row)
	}
	table.Render()
}

func errorHandle(e error) {
	if e != nil {
		panic(e)
	}
}

func Download(song platform.Song) {
	resp, e := http.Get(song.Url.Url)
	if e != nil {
		fmt.Println(song.Url.Url, e)
		return
	}

	bar := pb.StartNew(song.Url.Size)
	bar.SetWidth(80)
	tmpl := `{{ red "` + song.Title + `:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{percent .}}`
	bar.SetTemplateString(tmpl)

	reader := bufio.NewReader(resp.Body)
	errorHandle(e)

	usr, _ := user.Current()
	downloadDir := filepath.Join(usr.HomeDir, "Downloads")
	if _, e := os.Stat(downloadDir); e != nil {
		_ = os.Mkdir(downloadDir, os.ModePerm)
	}

	fileHandler, e := os.Create(filepath.Join(downloadDir, song.Singer+"-"+song.Title+".mp3"))
	errorHandle(e)
	for {
		p := make([]byte, 512)
		l, e := reader.Read(p)
		if e != nil && l == 0 {
			break
		}
		_, _ = fileHandler.Write(p[:l])
		bar.Add(l)
	}
	_ = fileHandler.Close()
	bar.Finish()
}
