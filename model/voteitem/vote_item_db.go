package voteitem

import (
	"context"
	"errors"
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "vote_items"

var (
	ErrAlreadyDeleted   = errors.New("already deleted votes")
	ErrNotExistVoteItem = errors.New("not exist votes")
	ErrInvalidDataType  = errors.New("invalid data type")
)

func Insert(o VoteItem) error {
	_, err := db.GetCollection(CollectionName).InsertOne(context.TODO(), o)
	return err
}

func Delete(id string) error {
	_, err := db.GetCollection(CollectionName).
		DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	return nil
}

func Update(o VoteItem) error {
	_, err := db.GetCollection(CollectionName).
		UpdateOne(
			context.TODO(),
			bson.M{"_id": o.Id},
			bson.D{{Key: "$set", Value: o}},
		)
	return err
}

func FindOne(filter bson.D) (VoteItem, error) {
	var o VoteItem
	result := db.GetCollection(CollectionName).
		FindOne(context.TODO(), filter)

	if result.Err() != nil {
		fmt.Println(result.Err().Error())
		if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
			return o, ErrNotExistVoteItem
		}
	}

	if err := result.Decode(&o); err != nil {
		fmt.Println(err.Error())
		return o, ErrInvalidDataType
	}

	return o, nil
}

func Find(filter bson.D) ([]VoteItem, error) {
	cur, err := db.GetCollection(CollectionName).
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var voteItems []VoteItem
	for cur.Next(context.TODO()) {
		var o VoteItem
		err := cur.Decode(&o)
		if err != nil {
			fmt.Println(err)
			continue
		}

		voteItems = append(voteItems, o)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}
	return voteItems, nil
}

func GetVoteItem(o VoteItem) (VoteItem, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", o.Id}},
		}},
	}
	return FindOne(filter)
}

func GetVoteItemByVoteId(voteId string) ([]VoteItem, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"voteId", voteId}},
		}},
	}
	return Find(filter)
}

func GetVoteItemAll() ([]VoteItem, error) {
	filter := bson.D{}
	return Find(filter)
}
