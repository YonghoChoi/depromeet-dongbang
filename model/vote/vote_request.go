package vote

type VoteRequest struct {
	VoteCommon `bson:",inline"`
	VoteItems  []string `json:"voteItems"`
}

func (o *VoteRequest) SetByVote(v VoteCommon) {
	o.VoteCommon = v
}
