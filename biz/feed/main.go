package main

import (
	feed "github.com/HUST-MiniTiktok/mini_tiktok/service/feed/kitex_gen/feed/feedservice"
	"log"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
