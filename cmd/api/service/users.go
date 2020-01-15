package service

import (
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers() ([]user.User, error) {
	return user.Find(bson.D{})
}
