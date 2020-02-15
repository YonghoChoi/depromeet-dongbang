package vote

type VoteRequest struct {
	VoteCommon
	VoteItems []string `json:"voteItems"`
}

func (o *VoteRequest) UpdateByVote(v VoteCommon) {
	o.VoteCommon = v
}
