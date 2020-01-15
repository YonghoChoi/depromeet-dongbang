package service

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(reqUser user.User) (user.User, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", reqUser.Id}},
			bson.D{{"name", reqUser.Name}},
		}},
	}
	findUser, err := user.FindOne(filter)
	if err != nil {
		return findUser, err
	}

	if reqUser.Token != findUser.Token {
		return reqUser, fmt.Errorf("invalid token")
	}

	return findUser, nil
}
