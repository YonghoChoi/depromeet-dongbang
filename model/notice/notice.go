package notice

import (
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/google/uuid"
	"time"
)

type Notice struct {
	Id         string    `json:"id" bson:"_id"`
	Pin        bool      `json:"pin" bson:"pin"`
	Category   string    `json:"category" bson:"category"`
	User       user.User `json:"user" bson:"user"`
	Title      string    `json:"title" bson:"title"`
	Content    string    `json:"content" bson:"content"`
	Images     []string  `json:"images" bson:"images"`
	CreateTime time.Time `json:"createTime" bson:"createTime"`
	UpdateTime time.Time `json:"updateTime" bson:"updateTime"`
}

func (o *Notice) Update(n Notice) {
	o.Pin = n.Pin
	o.Category = n.Category
	o.Title = n.Title
	o.Content = n.Content
	o.Images = n.Images
	o.UpdateTime = time.Now()
}

func New(u user.User, title, content, category string, images []string) Notice {
	now := time.Now()
	n := Notice{}
	n.Id = uuid.New().String()
	n.User = u
	n.Title = title
	n.Content = content
	n.Category = category
	n.Images = images
	n.CreateTime = now
	n.UpdateTime = now
	return n
}
