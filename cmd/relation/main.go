package main

import (
	relation "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/relation/relationservice"
	"log"
)

func main() {
	svr := relation.NewServer(new(RelationServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
