package main

import (
	user "github.com/HUST-MiniTiktok/mini_tiktok/service/user/kitex_gen/user/userservice"
	"log"
)

func main() {
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
