package user

import (
	"context"
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "users"

func Insert(u User) error {
	_, err := db.GetCollection(CollectionName).InsertOne(context.TODO(), u)
	return err
}

func Delete(u User) error {
	_, err := db.GetCollection(CollectionName).
		DeleteOne(context.TODO(), bson.M{"_id": u.Id})
	return err
}

func Update(u User) error {
	_, err := db.GetCollection(CollectionName).
		UpdateOne(
			context.TODO(),
			bson.M{"_id": u.Id},
			bson.D{{Key: "$set", Value: u}},
		)
	return err
}

func FindOne(filter bson.D) (User, error) {
	var u User
	result := db.GetCollection(CollectionName).
		FindOne(context.TODO(), filter)

	if result.Err() != nil {
		fmt.Println(result.Err().Error())
		if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
			return u, fmt.Errorf("not exist user")
		}
	}

	if err := result.Decode(&u); err != nil {
		fmt.Println(err.Error())
		return u, fmt.Errorf("invalid user data type")
	}

	return u, nil
}

func Find(filter bson.D) ([]User, error) {
	cur, err := db.GetCollection(CollectionName).
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var users []User
	for cur.Next(context.TODO()) {
		var u User
		err := cur.Decode(&u)
		if err != nil {
			fmt.Println(err)
			continue
		}

		users = append(users, u)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}
	return users, nil
}
