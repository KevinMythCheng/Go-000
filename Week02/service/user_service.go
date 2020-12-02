package service

import (
	"github.com/kevinlcheng/gotrain/Week02/dao"
	"github.com/kevinlcheng/gotrain/Week02/entity"
	"github.com/pkg/errors"
	"sync"
)

type userService struct {
}

var (
	service  *userService
	userOnce sync.Once
)

func Build() *userService {
	userOnce.Do(func() {
		service = &userService{}
	})
	return service
}

func (s *userService) FindUser(id int64) (*entity.User, error) {
	_, err := dao.Build().FindUser(id)
	if err != nil {
		if errors.Is(err, dao.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	if id == 1 {
		return nil, errors.Wrap(errors.New("Can not set as No.1 member"), "error to find the User")
	}

	return nil, nil
}
