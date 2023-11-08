package db

import (
	"context"
	"encoding/json"

	"gorm.io/gorm"

	"github.com/HUST-MiniTiktok/mini_tiktok/pkg/mw/mq"
)

type FavoriteMessage struct {
	ActionType int32
	UserId     int64
	VideoId    int64
}

func LoadSyncMQ() {
	ctx := context.Background()
	go FavoriteSync(ctx)
}

func FavoriteSync(ctx context.Context) {
	Sync := new(SyncFavorite)
	err := Sync.SyncFavorite(ctx)
	if err != nil {
		return
	}
}

type SyncFavorite struct {
}

func (s *SyncFavorite) SyncFavorite(ctx context.Context) error {
	msg, err := mq.ConsumeMessage(ctx, FavoriteMQName)
	if err != nil {
		return err
	}
	var forever chan struct{}
	go func() {
		for d := range msg {
			var req *FavoriteMessage
			err = json.Unmarshal(d.Body, &req)
			if err != nil {
				return
			}
			err = FavoriteMQ2DB(ctx, req)
			if err != nil {
				return
			}
			d.Ack(false)
		}
	}()
	<-forever
	return nil
}

func FavoriteMQ2DB(ctx context.Context, message *FavoriteMessage) error {

	if message.ActionType == 1 {
		favorite := Favorite{
			UserId:  message.UserId,
			VideoId: message.VideoId}

		err := DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
			if err := tx.WithContext(ctx).Model(&Favorite{}).Create(&favorite).Error; err != nil {
				return err
			}
			return nil
		})

		if err != nil {
			return err
		}

	} else if message.ActionType == 2 {
		err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ? and video_id = ?", message.UserId, message.VideoId).Delete(&Favorite{}).Error
		if err != nil {
			return err
		}
	}

	return nil
}
