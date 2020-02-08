package user

import (
	"flag"
	"github.com/YonghoChoi/depromeet-dongbang/cmd/api/conf"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func setup() {
	configPath := flag.String("config", "../../cmd/api/conf/config.yml", "Input config file path")
	flag.Parse()
	conf.SetConfigFilePath(*configPath)
}

func TestInsert(t *testing.T) {
	setup()
	u := New("test user", "test token", "")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}
}

func TestUpdate(t *testing.T) {
	setup()
	u := New("test user", "test token")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}

	u.Name = "modify user"
	if err := Update(u); err != nil {
		t.Fatal(err.Error())
	}

	t.Log(u)
}

func TestDelete(t *testing.T) {
	setup()
	u := New("test user", "test token")
	if err := Insert(u); err != nil {
		t.Fatal(err.Error())
	}

	if err := Delete(u); err != nil {
		t.Fatal(err.Error())
	}

	t.Log(u)
}

func TestFindOne(t *testing.T) {
	setup()
	id := "af751898-4418-4940-b64f-6413a71e75b6"
	filter := bson.D{
		{"_id", id},
	}
	u, err := FindOne(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(u)
}

func TestFind(t *testing.T) {
	setup()
	id := "af751898-4418-4940-b64f-6413a71e75b6"
	filter := bson.D{
		{"$or", []interface{}{
			bson.D{{"_id", id}},
			bson.D{{"name", "test user"}},
		}},
	}
	users, err := Find(filter)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Log(users)
}
