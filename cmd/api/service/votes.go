package service

import (
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/YonghoChoi/depromeet-dongbang/model/vote"
	"github.com/YonghoChoi/depromeet-dongbang/model/voteitem"
	"github.com/YonghoChoi/depromeet-dongbang/model/voter"
)

func GetVotes() ([]vote.VoteResponse, error) {
	votes, err := vote.GetVoteAll()
	if err != nil {
		return nil, err
	}

	var result []vote.VoteResponse
	for _, v := range votes {
		resp := vote.VoteResponse{
			VoteCommon: v.VoteCommon,
		}

		// 사용자 정보 갱신
		u, err := user.GetUser(user.User{Id: v.WriterId})
		if err != nil {
			return nil, err
		}
		resp.Writer = u

		// 투표 항목 로드
		findVoteItems, err := voteitem.GetVoteItemByVoteId(v.Id)
		if err != nil {
			return nil, err
		}

		// 각 투표 항목 정보 수집
		for _, item := range findVoteItems {
			findVoters, err := voter.GetVoterByVoteItemId(item.Id)
			if err != nil {
				return nil, err
			}

			resp.VoteItems = append(resp.VoteItems, vote.VoteItem{
				Id:        item.Id,
				Content:   item.Content,
				UserCount: len(findVoters),
				Order:     item.Order,
			})
		}

		result = append(result, resp)
	}

	return result, nil
}

func CreateVote(voteReq vote.VoteRequest) (vote.VoteResponse, error) {
	voteResp := vote.VoteResponse{}
	writer, err := user.GetUser(user.User{Id: voteReq.WriterId})
	if err != nil {
		return vote.VoteResponse{}, err
	}

	v := vote.New(voteReq.WriterId, voteReq.Title, voteReq.Content, voteReq.Options, voteReq.ClosingTime)
	if err := vote.Insert(v); err != nil {
		return vote.VoteResponse{}, err
	}

	voteReq.Id = v.Id
	voteResp.UpdateByVote(v.VoteCommon)
	voteResp.Writer = writer
	for i, content := range voteReq.VoteItems {
		voteItem := voteitem.New(voteReq.Id, content)
		if err := voteitem.Insert(voteItem); err != nil {
			return vote.VoteResponse{}, err
		}

		voteResp.VoteItems = append(voteResp.VoteItems, vote.VoteItem{
			Id:        voteItem.Id,
			Content:   voteItem.Content,
			UserCount: 0,
			Order:     i,
		})
	}

	return voteResp, nil
}

func EditVote(id string, v vote.Vote) (result vote.Vote, err error) {
	result, err = vote.GetVote(vote.Vote{VoteCommon: vote.VoteCommon{Id: id}})
	if err != nil {
		return
	}

	result.Update(v)
	return
}

func DelVote(id string) (vote.VoteResponse, error) {
	result := vote.VoteResponse{}
	v, err := vote.GetVote(vote.Vote{
		VoteCommon: vote.VoteCommon{Id: id},
	})

	if err != nil {
		if err == vote.ErrNotExistVote {
			err = vote.ErrAlreadyDeleted
			return result, err
		}

		return result, err
	}
	result.UpdateByVote(v.VoteCommon)

	voteItems, err := voteitem.GetVoteItemByVoteId(id)
	if err != nil {
		return result, err
	}

	for _, item := range voteItems {
		voters, err := voter.GetVoterByVoteItemId(item.Id)
		if err != nil {
			return result, err
		}

		for _, vt := range voters {
			if err := voter.Delete(vt); err != nil {
				return result, err
			}
		}

		if err := voteitem.Delete(item); err != nil {
			return result, err
		}
	}

	if err := vote.Delete(v); err != nil {
		return result, err
	}

	return result, nil
}
