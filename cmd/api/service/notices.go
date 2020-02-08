package service

import (
	"fmt"
	"github.com/YonghoChoi/depromeet-dongbang/model/notice"
	"github.com/YonghoChoi/depromeet-dongbang/model/user"
)

func GetNotices() ([]notice.Notice, error) {
	notices, err := notice.GetNoticeAll()
	if err != nil {
		return nil, err
	}

	for i := range notices {
		u, err := user.GetUser(notices[i].User)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		notices[i].User = u
	}
	return notices, nil
}

func CreateNotice(n notice.Notice) (notice.Notice, error) {
	u, err := user.GetUser(n.User)
	if err != nil {
		return notice.Notice{}, err
	}

	n = notice.New(u, n.Title, n.Content, n.Category, n.Images)
	if err := notice.Insert(n); err != nil {
		return notice.Notice{}, err
	}
	return n, nil
}

func EditNotice(id string, n notice.Notice) (notice.Notice, error) {
	findNotice, err := notice.GetNotice(notice.Notice{Id: id})
	if err != nil {
		return notice.Notice{}, err
	}
	findNotice.Update(n)
	return findNotice, notice.Update(n)
}

func DelNotice(id string) (notice.Notice, error) {
	noti, err := notice.GetNotice(notice.Notice{Id: id})
	if err != nil {
		if err == notice.ErrNotExistNotice {
			return noti, notice.ErrAlreadyDeleted
		}

		return noti, err
	}
	return noti, notice.Delete(notice.Notice{Id: id})
}
