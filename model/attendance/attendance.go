package attendance

import (
	"github.com/google/uuid"
	"time"
)

type Attendance struct {
	Id         string    `json:"id" bson:"_id"`
	Token      string    `json:"token" bson:"token"`
	ExpireTime time.Time `json:"expireTime" bson:"expireTime"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
}

func (o *Attendance) Update(arg Attendance) {
	o.Token = arg.Token
	o.ExpireTime = arg.ExpireTime
	o.UpdateTime = time.Now()
}

func New(token string, expireTime time.Time) Attendance {
	now := time.Now()
	o := Attendance{}
	o.Id = uuid.New().String()
	o.Token = token
	o.ExpireTime = expireTime
	o.CreateTime = now
	o.UpdateTime = now
	return o
}
