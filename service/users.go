package service

import (
	"mixindev/model"
)

func ListUsers(offset, limit int) ([]model.UserInfo, int64, error) {
	var infos []model.UserInfo
	var u model.User
	users, count, err := u.ListUser(offset, limit)
	if err != nil {
		return nil, count, err
	}

	for _, user := range users {
		infos = append(infos, model.UserInfo{
			Id:       int(user.ID),
			Username: user.Username,
			Avatar:   user.Avatar,
		})
	}

	return infos, count, nil
}
