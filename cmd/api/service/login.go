package service

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
)

func Login(reqUser user.User) (user.User, error) {
	findUser, err := user.GetUser(reqUser)
	if err != nil {
		return findUser, err
	}

	if reqUser.Token != findUser.Token {
		return reqUser, fmt.Errorf("invalid token")
	}

	return findUser, nil
}
