package service

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
)

func Join(u user.User) (user.User, error) {
	findUser, err := user.GetUser(u)
	if err != nil {
		if err != user.ErrNotExistUser {
			fmt.Println(err.Error())
			return u, fmt.Errorf("database error.\n")
		}
	}

	if findUser.Id != "" {
		return u, fmt.Errorf("already exist user")
	}

	u.Join()
	if err := user.Insert(u); err != nil {
		fmt.Println(err.Error())
		return u, fmt.Errorf("database error.\n")
	}

	fmt.Printf("join sucess. %v\n", u)
	return u, nil
}
