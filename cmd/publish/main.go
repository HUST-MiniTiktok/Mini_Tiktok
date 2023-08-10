package main

import (
	publish "github.com/HUST-MiniTiktok/mini_tiktok/cmd/publish/kitex_gen/publish/publishservice"
	"log"
)

func main() {
	svr := publish.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
