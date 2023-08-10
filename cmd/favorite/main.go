package main

import (
	favorite "github.com/HUST-MiniTiktok/mini_tiktok/cmd/favorite/kitex_gen/favorite/favoriteservice"
	"log"
)

func main() {
	svr := favorite.NewServer(new(FavoriteServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
