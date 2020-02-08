package vote

import (
	"context"
	"errors"
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "votes"

var (
	ErrAlreadyDeleted  = errors.New("already deleted votes")
	ErrNotExistVote    = errors.New("not exist votes")
	ErrInvalidDataType = errors.New("invalid data type")
)

func Insert(o Vote) error {
	_, err := db.GetCollection(CollectionName).InsertOne(context.TODO(), o)
	return err
}

func Delete(o Vote) error {
	_, err := db.GetCollection(CollectionName).
		DeleteOne(context.TODO(), bson.M{"_id": o.Id})
	if err != nil {
		return err
	}

	return nil
}

func Update(o Vote) error {
	_, err := db.GetCollection(CollectionName).
		UpdateOne(
			context.TODO(),
			bson.M{"_id": o.Id},
			bson.D{{Key: "$set", Value: o}},
		)
	return err
}

func FindOne(filter bson.D) (Vote, error) {
	var o Vote
	result := db.GetCollection(CollectionName).
		FindOne(context.TODO(), filter)

	if result.Err() != nil {
		fmt.Println(result.Err().Error())
		if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
			return o, ErrNotExistVote
		}
	}

	if err := result.Decode(&o); err != nil {
		fmt.Println(err.Error())
		return o, ErrInvalidDataType
	}

	return o, nil
}

func Find(filter bson.D) ([]Vote, error) {
	cur, err := db.GetCollection(CollectionName).
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var notices []Vote
	for cur.Next(context.TODO()) {
		var o Vote
		err := cur.Decode(&o)
		if err != nil {
			fmt.Println(err)
			continue
		}

		notices = append(notices, o)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}
	return notices, nil
}

func GetVote(o Vote) (Vote, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", o.Id}},
		}},
	}
	return FindOne(filter)
}

func GetVoteAll() ([]Vote, error) {
	filter := bson.D{}
	return Find(filter)
}
