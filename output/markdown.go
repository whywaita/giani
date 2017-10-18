package output

import (
	"fmt"

	"github.com/whywaita/gigani/parse"
)

func Markdown(animes []parse.Anime) {
	// output Markdown table
	fmt.Println(`# 録画リスト`)
	fmt.Println(`| 名前 | 局 | 放送時間 | done |`)
	fmt.Println(`|---|---|---|---|`)

	for _, anime := range animes {
		fmt.Println(`|[` + anime.Name + `](` + anime.URL + `) |` + anime.BloadCaster + `|` + anime.StartDate + `||`)
	}
}