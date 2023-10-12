package db

import (
	"context"
	"log"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Init()
}

func TestNewComment(t *testing.T) {
	ctx := context.Background()
	Init()
	_, err := NewComment(ctx, 1, 1, "一眼丁真")
	log.Println(err)
}

func TestDelComment(t *testing.T) {
	ctx := context.Background()
	Init()
	newcomment, _ := NewComment(ctx, 1, 1, "一眼丁真")
	err := DelComment(ctx, newcomment.ID, 1)
	log.Println(err)
}

func TestGetVideoComments(t *testing.T) {
	ctx := context.Background()
	Init()
	newcomment1, _ := NewComment(ctx, 1, 1, "亿烟丁真")
	newcomment2, _ := NewComment(ctx, 2, 1, "鉴定为 王源剩太多导致的")
	comments, err := GetVideoComments(ctx, 1)
	log.Println(comments[1].CommentText)
	log.Println(err)
	DelComment(ctx, newcomment1.ID, 1)
	comments1, err1 := GetVideoComments(ctx, 1)
	log.Println(comments1[0].CommentText)
	log.Println(err1)
	DelComment(ctx, newcomment2.ID, 1)
	comments2, err2 := GetVideoComments(ctx, 1)
	log.Println(comments2[0].CommentText)
	log.Println(err2)
}

func TestGetVideoCommentCounts(t *testing.T) {
	ctx := context.Background()
	Init()

	counts, err := GetVideoCommentCount(ctx, 1)
	log.Println(counts)
	log.Println(err)

	newcomment1, _ := NewComment(ctx, 1, 1, "亿烟丁真")
	newcomment2, _ := NewComment(ctx, 2, 1, "鉴定为 王源剩太多导致的")
	time.Sleep(time.Second)
	counts, err = GetVideoCommentCount(ctx, 1)
	log.Println(counts)
	log.Println(err)

	DelComment(ctx, newcomment1.ID, 1)
	time.Sleep(time.Second)
	counts, err = GetVideoCommentCount(ctx, 1)
	log.Println(counts)
	log.Println(err)

	DelComment(ctx, newcomment2.ID, 1)
	time.Sleep(time.Second)
	counts, err = GetVideoCommentCount(ctx, 1)
	log.Println(counts)
	log.Println(err)
}
