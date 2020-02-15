package voteitem

import (
	"github.com/google/uuid"
	"time"
)

type VoteItem struct {
	Id         string    `json:"id" bson:"_id"`
	VoteId     string    `json:"voteId" bson:"voteId"`
	Content    string    `json:"content" bson:"content"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
}

func (o *VoteItem) Update(arg VoteItem) {
	o.Content = arg.Content
	o.UpdateTime = time.Now()
}

func New(voteId, content string) VoteItem {
	now := time.Now()
	n := VoteItem{}
	n.Id = uuid.New().String()
	n.VoteId = voteId
	n.Content = content
	n.CreateTime = now
	n.UpdateTime = now
	return n
}
