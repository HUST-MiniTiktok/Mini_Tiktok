package pack

import (
	"context"

	"github.com/HUST-MiniTiktok/mini_tiktok/cmd/message/client"
	user "github.com/HUST-MiniTiktok/mini_tiktok/kitex_gen/user"
)

func IsExistUser(ctx context.Context, user_id int64) (bool, error) {
	resp, err := client.UserRPC.CheckUserIsExist(ctx, &user.CheckUserIsExistRequest{UserId: user_id})
	if err != nil {
		return false, err
	}
	return resp.IsExist, nil
}
