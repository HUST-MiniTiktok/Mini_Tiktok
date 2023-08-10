package main

import (
	message "github.com/HUST-MiniTiktok/mini_tiktok/biz/message/kitex_gen/message/messageservice"
	"log"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
