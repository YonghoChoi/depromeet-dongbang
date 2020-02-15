package vote

import (
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

type VoteCommon struct {
	Id          string     `json:"id" bson:"_id"`
	WriterId    string     `json:"writerId" bson:"writerId"`
	VoteStatus  VoteStatus `json:"voteStatus" bson:"voteStatus"`
	Title       string     `json:"title" bson:"title"`
	Content     string     `json:"content" bson:"content"`
	Options     []Option   `json:"options" bson:"options"`
	ClosingTime time.Time  `json:"closingTime" bson:"closingTime"`
}
