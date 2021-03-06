package voteitem

import (
	"github.com/google/uuid"
	"time"
)

type VoteItem struct {
	Id         string    `json:"id" bson:"_id"`
	VoteId     string    `json:"voteId" bson:"voteId"`
	Content    string    `json:"content" bson:"content"`
	Order      int       `json:"order" bson:"order"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
}

func (o *VoteItem) Update(arg VoteItem) {
	o.Content = arg.Content
	o.UpdateTime = time.Now()
}

func New(voteId, content string, order int) VoteItem {
	now := time.Now()
	n := VoteItem{}
	n.Id = uuid.New().String()
	n.VoteId = voteId
	n.Content = content
	n.Order = order
	n.CreateTime = now
	n.UpdateTime = now
	return n
}
