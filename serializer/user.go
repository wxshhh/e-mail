package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

// User VO：传给前端展示的数据类型
type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	NickName  string `json:"nick_name"`
	Type      int    `json:"type"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	CreatedAt int64  `json:"create_at"`
}

func BuildUser(user *model.User) *User {
	return &User{
		ID:        user.ID,
		Username:  user.Username,
		NickName:  user.NickName,
		Email:     user.Email,
		Status:    user.Status,
		Avatar:    conf.Host + conf.HttpPort + conf.AvatarPath + user.Avatar,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
