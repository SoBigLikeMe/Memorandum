package serialzer

import "memorandum/model"

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`                 // 用户id
	UserName string `json:"user_name" from:"user_name" example:"yfc"` // 用户名
	Status   string `json:"status" from:"status"`                     // 用户状态
	CreateAt int64  `json:"create_at" from:"create_at"`
}

//BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}
