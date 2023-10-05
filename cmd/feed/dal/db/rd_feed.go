package db

// import (
// 	"encoding/json"
// 	"math/rand"
// 	"strconv"
// 	"time"

// 	common "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/common"
// 	"github.com/cloudwego/kitex/pkg/klog"
// )

// const (
// 	FavoriteCountField   = "favoriteCount"
// 	CommentCountField    = "commentCount"
// 	VideoInfoSuffix      = ":videoInfo"
// 	VideoInfoCountSuffix = ":videoInfoCount"
// 	RDCacheExpire        = time.Hour
// )

// type VideoInfo struct {
// 	// Author
// 	AuthorId      int64
// 	Name          string
// 	FollowCount   int64
// 	FollowerCount int64
// 	// IsFollow        bool
// 	Avatar          string
// 	BackgroundImage string
// 	Signature       string
// 	TotalFavorited  int64
// 	WorkCount       int64
// 	FavoriteCount   int64
// 	// video
// 	PlayURL  string
// 	CoverURL string
// 	// IsFavorite bool
// 	Title string
// }

// // type VideoInfoCount struct {
// // 	FavoriteCount int64
// // 	CommentCount  int64
// // }

// func RDExistVideo(video_id int64) bool {
// 	video_id_str := strconv.FormatInt(video_id, 10) + VideoInfoSuffix
// 	return RDClient.Exists(video_id_str)
// }

// func RDGetVideoInfo(video_id int64) string {
// 	video_id_str := strconv.FormatInt(video_id, 10) + VideoInfoSuffix
// 	return RDClient.Get(video_id_str)
// }

// func RDSetVideoInfo(video_id int64, common_video *common.Video) error {
// 	video_info_key := strconv.FormatInt(video_id, 10) + VideoInfoSuffix
// 	video_info := &VideoInfo{
// 		AuthorId:        common_video.Author.Id,
// 		Name:            common_video.Author.Name,
// 		FollowCount:     *common_video.Author.FollowCount,
// 		FollowerCount:   *common_video.Author.FollowerCount,
// 		Avatar:          *common_video.Author.Avatar,
// 		BackgroundImage: *common_video.Author.BackgroundImage,
// 		Signature:       *common_video.Author.Signature,
// 		TotalFavorited:  *common_video.Author.TotalFavorited,
// 		WorkCount:       *common_video.Author.WorkCount,
// 		FavoriteCount:   *common_video.Author.FavoriteCount,
// 		PlayURL:         common_video.PlayUrl,
// 		CoverURL:        common_video.CoverUrl,
// 		Title:           common_video.Title,
// 	}

// 	// JSON serialization
// 	video_info_json, err := json.Marshal(video_info)
// 	if err != nil {
// 		klog.Errorf("json Marshal video_info failed: %v", err.Error())
// 		return err
// 	}

// 	// open pipeline
// 	pipeline := RDClient.Client.Pipeline()

// 	// set VideoInfo JSON cache
// 	err = pipeline.Set(video_info_key, video_info_json,
// 		RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute).Err()
// 	if err != nil {
// 		klog.Errorf("set VideoInfo JSON cache failed: %v", err.Error())
// 		return err
// 	}

// 	// video_Info_Count := &VideoInfoCount{
// 	// 	FavoriteCount: common_video.FavoriteCount,
// 	// 	CommentCount:  common_video.CommentCount,
// 	// }

// 	info_count_key := strconv.FormatInt(video_id, 10) + VideoInfoCountSuffix

// 	// 使用 MSet 进行批量设置
// 	err = pipeline.HMSet(info_count_key, map[string]interface{}{
// 		FavoriteCountField: common_video.FavoriteCount,
// 		CommentCountField:  common_video.CommentCount,
// 	}).Err()
// 	if err != nil {
// 		klog.Error("redis pipeline set VideoInfoCount cache failed: %v", err.Error())
// 		return err
// 	}
// 	err = pipeline.Expire(info_count_key, RDCacheExpire+time.Duration(rand.Intn(200))*time.Minute).Err()
// 	if err != nil {
// 		klog.Error("redis pipeline set VideoInfoCount Expire failed: %v", err.Error())
// 		return err
// 	}
// 	// 执行管道中的命令
// 	_, err = pipeline.Exec()
// 	if err != nil {
// 		klog.Error("redis pipeline failed: %v", err.Error())
// 		return err
// 	}
// 	return nil
// }
