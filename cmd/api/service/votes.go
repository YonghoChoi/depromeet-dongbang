package service

import (
	"errors"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/YonghoChoi/depromeet-dongbang/model/vote"
	"github.com/YonghoChoi/depromeet-dongbang/model/voteitem"
	"github.com/YonghoChoi/depromeet-dongbang/model/voter"
)

// CreateVote : 투표 생성
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
	voteResp.VoteCommon = v.VoteCommon
	voteResp.Writer = writer
	voteResp.VoteItems, err = createVoteItems(v.Id, voteReq.VoteItems)
	if err != nil {
		return vote.VoteResponse{}, err
	}
	return voteResp, nil
}

// GetVotes : 투표 리스트 조회
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

			resp.VoteItems = append(resp.VoteItems, vote.VoteItemResp{
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

// GetVote : 단일 투표 조회
func GetVote(id string) (vote.VoteResponse, error) {
	v, err := vote.GetVote(vote.Vote{VoteCommon: vote.VoteCommon{Id: id}})
	if err != nil {
		return vote.VoteResponse{}, err
	}

	resp := vote.VoteResponse{
		VoteCommon: v.VoteCommon,
	}

	// 사용자 정보 갱신
	u, err := user.GetUser(user.User{Id: v.WriterId})
	if err != nil {
		return vote.VoteResponse{}, err
	}
	resp.Writer = u

	// 투표 항목 로드
	findVoteItems, err := voteitem.GetVoteItemByVoteId(v.Id)
	if err != nil {
		return vote.VoteResponse{}, err
	}

	// 각 투표 항목 정보 수집
	for _, item := range findVoteItems {
		findVoters, err := voter.GetVoterByVoteItemId(item.Id)
		if err != nil {
			return vote.VoteResponse{}, err
		}

		resp.VoteItems = append(resp.VoteItems, vote.VoteItemResp{
			Id:        item.Id,
			Content:   item.Content,
			UserCount: len(findVoters),
			Order:     item.Order,
		})
	}
	return resp, nil
}

// EditVote : 투표 수정
func EditVote(voteReq vote.VoteRequest) (vote.VoteResponse, error) {
	voteResp := vote.VoteResponse{}
	writer, err := user.GetUser(user.User{Id: voteReq.WriterId})
	if err != nil {
		return voteResp, err
	}

	vt, err := GetVote(voteReq.Id)
	if err != nil {
		return voteResp, err
	}

	if vt.HasVote() {
		return voteResp, errors.New("already vote some user")
	}

	for _, item := range vt.VoteItems {
		if err := voteitem.Delete(item.Id); err != nil {
			return voteResp, err
		}
	}

	voteReq.Id = vt.Id
	voteResp.VoteCommon = vt.VoteCommon
	voteResp.Writer = writer
	voteResp.VoteItems, err = createVoteItems(vt.Id, voteReq.VoteItems)
	if err != nil {
		return voteResp, err
	}

	return voteResp, nil
}

// DelVote : 단일 투표 제거
func DelVote(id string) (vote.VoteResponse, error) {
	result := vote.VoteResponse{}
	v, err := vote.GetVote(vote.Vote{
		VoteCommon: vote.VoteCommon{Id: id},
	})

	if err != nil {
		if err == vote.ErrNotExistVote {
			return result, vote.ErrAlreadyDeleted
		}

		return result, err
	}

	result.VoteCommon = v.VoteCommon
	if err := delVoteItemsByVoteId(id); err != nil {
		return result, err
	}

	if err := vote.Delete(v); err != nil {
		return result, err
	}

	return result, nil
}

func delVoteItemsByVoteId(voteId string) error {
	voteItems, err := voteitem.GetVoteItemByVoteId(voteId)
	if err != nil {
		return err
	}

	for _, item := range voteItems {
		voters, err := voter.GetVoterByVoteItemId(item.Id)
		if err != nil {
			return err
		}

		for _, vt := range voters {
			if err := voter.Delete(vt); err != nil {
				return err
			}
		}

		if err := voteitem.Delete(item.Id); err != nil {
			return err
		}
	}

	return nil
}

func createVoteItems(voteId string, contents []string) ([]vote.VoteItemResp, error) {
	var voteItems []vote.VoteItemResp
	for i, content := range contents {
		voteItem := voteitem.New(voteId, content, i)
		if err := voteitem.Insert(voteItem); err != nil {
			return nil, err
		}

		voteItems = append(voteItems, vote.VoteItemResp{
			Id:        voteItem.Id,
			Content:   voteItem.Content,
			UserCount: 0,
			Order:     voteItem.Order,
		})
	}

	return voteItems, nil
}
