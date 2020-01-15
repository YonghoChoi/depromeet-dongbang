package service

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"go.mongodb.org/mongo-driver/bson"
)

func Join(u user.User) (user.User, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", u.Id}},
			bson.D{{"name", u.Name}},
		}},
	}

	users, err := user.Find(filter)
	if err != nil {
		fmt.Println(err.Error())
		return u, fmt.Errorf("database error.\n")
	}

	if len(users) > 0 {
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
