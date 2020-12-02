package dao

import (
	"fmt"
	"github.com/pkg/errors"
	"sync"
)

type userDao struct {
	name string
	age  int64
}

var (
	ErrNoRows   = errors.New("No rows found")
	userDaoOnce sync.Once
	dao         *userDao
)

func Build() *userDao {
	userDaoOnce.Do(func() {
		dao = &userDao{}
	})
	return dao
}

func (u *userDao) FindUser(id int64) (int64, error) {
	if id < 0 {
		return 0, errors.Wrap(ErrNoRows, fmt.Sprintf("no rows found use id is {%d}", id))
	}

	return 0, nil
}
