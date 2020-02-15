package vote

import "github.com/YonghoChoi/depromeet-dongbang/model/user"

type VoteResponse struct {
	VoteCommon
	Writer    user.User  `json:"writer"`
	VoteItems []VoteItem `json:"voteItems"`
}

func (o *VoteResponse) UpdateByVote(v VoteCommon) {
	o.VoteCommon = v
}

type VoteItem struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	UserCount int    `json:"userCount"`
	Order     int    `json:"order"`
}
