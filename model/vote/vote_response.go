package vote

import "github.com/YonghoChoi/depromeet-dongbang/model/user"

type VoteResponse struct {
	VoteCommon `bson:",inline"`
	Writer     user.User      `json:"writer"`
	VoteItems  []VoteItemResp `json:"voteItems"`
}

func (o *VoteResponse) SetByVote(v VoteCommon) {
	o.VoteCommon = v
}

func (o *VoteResponse) HasVote() bool {
	for _, item := range o.VoteItems {
		if item.UserCount > 0 {
			return true
		}
	}

	return false
}

type VoteItemResp struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	UserCount int    `json:"userCount"`
	Order     int    `json:"order"`
}
