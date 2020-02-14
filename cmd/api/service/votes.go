package service

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
	"github.com/YonghoChoi/depromeet-dongbang/model/vote"
)

func GetVotes() ([]vote.Vote, error) {
	votes, err := vote.GetVoteAll()
	if err != nil {
		return nil, err
	}

	for i := range votes {
		u, err := user.GetUser(votes[i].User)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		votes[i].User = u
	}
	return votes, nil
}

func CreateVote(v vote.Vote) (vote.Vote, error) {
	u, err := user.GetUser(v.User)
	if err != nil {
		return vote.Vote{}, err
	}

	v = vote.New(u, v.Title, v.Content, v.Votes, v.Options, v.ClosingTime)
	if err := vote.Insert(v); err != nil {
		return vote.Vote{}, err
	}
	return v, nil
}

func EditVote(id string, v vote.Vote) (result vote.Vote, err error) {
	result, err = vote.GetVote(vote.Vote{Id: id})
	if err != nil {
		return
	}

	result.Update(v)
	return
}

func DelVote(id string) (result vote.Vote, err error) {
	result, err = vote.GetVote(vote.Vote{Id: id})
	if err != nil {
		if err == vote.ErrNotExistVote {
			err = vote.ErrAlreadyDeleted
			return
		}

		return
	}

	err = vote.Delete(result)
	return
}
