package user

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id         string    `json:"id" bson:"_id"`
	Name       string    `json:"name"  bson:"name"`
	Token      string    `json:"token"  bson:"token"`
	ProfilePic string    `json:"profilePic"`
	CreateTime time.Time `json:"create_time"  bson:"create_time"`
	UpdateTime time.Time `json:"update_time"  bson:"update_time"`
}

func (o *User) Join() {
	o.Id = uuid.New().String()
	o.CreateTime = time.Now()
	o.UpdateTime = time.Now()
}

func New(name, token, profilePic string) User {
	now := time.Now()

	u := User{}
	u.Id = uuid.New().String()
	u.Name = name
	u.Token = token
	u.ProfilePic = profilePic
	u.CreateTime = now
	u.UpdateTime = now
	return u
}
