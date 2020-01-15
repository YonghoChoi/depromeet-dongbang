package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         string    `json:"id" bson:"_id"`
	Name       string    `json:"name"  bson:"name"`
	Token      string    `json:"token"  bson:"token"`
	CreateTime time.Time `json:"create_time"  bson:"create_time"`
	UpdateTime time.Time `json:"update_time"  bson:"update_time"`
}

func (o *User) Join() {
	o.Id = uuid.New().String()
	o.CreateTime = time.Now()
	o.UpdateTime = time.Now()
}

func New(name, token string) User {
	u := User{}
	u.Id = uuid.New().String()
	u.Name = name
	u.Token = token
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()
	return u
}
