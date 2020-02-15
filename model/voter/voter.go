package voter

import (
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/google/uuid"
	"time"
)

type Voter struct {
	Id         string    `json:"id" bson:"_id"`
	VoteItemId string    `json:"votePaperId" bson:"votePaperId"`
	User       user.User `json:"user" bson:"user"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
}

func (o *Voter) Update(arg Voter) {
	o.User = arg.User
	o.UpdateTime = time.Now()
}

func New(votePaperId string, u user.User) Voter {
	now := time.Now()
	n := Voter{}
	n.Id = uuid.New().String()
	n.VoteItemId = votePaperId
	n.User = u
	n.CreateTime = now
	n.UpdateTime = now
	return n
}
