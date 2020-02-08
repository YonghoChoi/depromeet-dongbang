package notice

import (
	"context"
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/pkg/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const CollectionName = "notices"

func Insert(o Notice) error {
	_, err := db.GetCollection(CollectionName).InsertOne(context.TODO(), o)
	return err
}

func Delete(o Notice) error {
	_, err := db.GetCollection(CollectionName).
		DeleteOne(context.TODO(), bson.M{"_id": o.Id})
	return err
}

func Update(o Notice) error {
	_, err := db.GetCollection(CollectionName).
		UpdateOne(
			context.TODO(),
			bson.M{"_id": o.Id},
			bson.D{{Key: "$set", Value: o}},
		)
	return err
}

func FindOne(filter bson.D) (Notice, error) {
	var o Notice
	result := db.GetCollection(CollectionName).
		FindOne(context.TODO(), filter)

	if result.Err() != nil {
		fmt.Println(result.Err().Error())
		if result.Err() == mongo.ErrNoDocuments || result.Err() == mongo.ErrNilDocument {
			return o, fmt.Errorf("not exist notice")
		}
	}

	if err := result.Decode(&o); err != nil {
		fmt.Println(err.Error())
		return o, fmt.Errorf("invalid notice data type")
	}

	return o, nil
}

func Find(filter bson.D) ([]Notice, error) {
	cur, err := db.GetCollection(CollectionName).
		Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	var notices []Notice
	for cur.Next(context.TODO()) {
		var o Notice
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

func GetNotice(o Notice) (Notice, error) {
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", o.Id}},
		}},
	}
	return FindOne(filter)
}

func GetNoticeAll() ([]Notice, error) {
	filter := bson.D{}
	return Find(filter)
}
