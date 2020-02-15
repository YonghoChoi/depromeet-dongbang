package vote

import (
	"github.com/google/uuid"
	"time"
)

type Vote struct {
	VoteCommon
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
}

func (o *Vote) Update(arg Vote) {
	o.VoteStatus = arg.VoteStatus
	o.Title = arg.Title
	o.Content = arg.Content
	o.Options = arg.Options
	o.ClosingTime = arg.ClosingTime
	o.UpdateTime = time.Now()
	if o.CreateTime.IsZero() {
		o.CreateTime = time.Now()
	}
}

func New(writerId, title, content string, options []Option, closingTime time.Time) Vote {
	now := time.Now()
	n := Vote{}
	n.Id = uuid.New().String()
	n.WriterId = writerId
	n.Title = title
	n.Content = content
	n.Options = options
	n.ClosingTime = closingTime
	n.CreateTime = now
	n.UpdateTime = now
	return n
}
