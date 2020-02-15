package vote

import (
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/google/uuid"
	"time"
)

type VoteStatus int
type Option int

const (
	Progress VoteStatus = 1 + iota
	Complete
)

var VoteStatuses = []string{
	"None",
	"progress",
	"complete",
}

const (
	Duplicate Option = 1 + iota
	Anonymous
	Date
)

var Options = []string{
	"None",
	"duplicate",
	"anonymous",
	"date",
}

func (o Option) String() string { return Options[o] }

type Vote struct {
	Id          string     `json:"id" bson:"_id"`
	User        user.User  `json:"user" bson:"user"`
	VoteStatus  VoteStatus `json:"voteStatus" bson:"voteStatus"`
	Title       string     `json:"title" bson:"title"`
	Content     string     `json:"content" bson:"content"`
	Options     []Option   `json:"options" bson:"options"`
	ClosingTime time.Time  `json:"closingTime" bson:"closingTime"`
	CreateTime  time.Time  `json:"createTime" bson:"createTime"`
	UpdateTime  time.Time  `json:"updateTime" bson:"updateTime"`
}

func (o *Vote) Update(arg Vote) {
	o.VoteStatus = arg.VoteStatus
	o.Title = arg.Title
	o.Content = arg.Content
	o.Options = arg.Options
	o.ClosingTime = arg.ClosingTime
	o.UpdateTime = time.Now()
}

func New(u user.User, title, content string, options []Option, closingTime time.Time) Vote {
	now := time.Now()
	n := Vote{}
	n.Id = uuid.New().String()
	n.User = u
	n.Title = title
	n.Content = content
	n.Options = options
	n.ClosingTime = closingTime
	n.CreateTime = now
	n.UpdateTime = now
	return n
}

type VoteMain struct {
	Vote
}
