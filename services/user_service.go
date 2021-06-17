package services

import (
	"go-skeleton/dao"
	"go-skeleton/model"
	"go-skeleton/model/constants"
	"go-skeleton/pkg/simpleDb"
)

var UserService = newUserTokenService()

func newUserService() *userService {
	return &userService{}
}

type userService struct {
}

func (s *userService) Get(id int64) *model.User {
	return dao.UserDao.Get(simpleDb.DB(), id)
}

func (s *userService) Take(where ...interface{}) *model.User {
	return dao.UserDao.Take(simpleDb.DB(), where...)
}

func (s *userService) Find(cnd *simpleDb.SqlCnd) []model.User {
	return dao.UserDao.Find(simpleDb.DB(), cnd)
}

func (s *userService) FindOne(cnd *simpleDb.SqlCnd) *model.User {
	return dao.UserDao.FindOne(simpleDb.DB(), cnd)
}

func (s *userService) FindPageByParams(params *simpleDb.QueryParams) (list []model.User, paging *simpleDb.Paging) {
	return dao.UserDao.FindPageByParams(simpleDb.DB(), params)
}

func (s *userService) FindPageByCnd(cnd *simpleDb.SqlCnd) (list []model.User, paging *simpleDb.Paging) {
	return dao.UserDao.FindPageByCnd(simpleDb.DB(), cnd)
}

func (s *userService) Update(t *model.User) error {
	err := dao.UserDao.Update(simpleDb.DB(), t)
	return err
}

func (s *userService) Updates(id int64, columns map[string]interface{}) error {
	err := dao.UserDao.Updates(simpleDb.DB(), id, columns)
	return err
}

func (s *userService) UpdateColumn(id int64, name string, value interface{}) error {
	err := dao.UserDao.UpdateColumn(simpleDb.DB(), id, name, value)
	return err
}

func (s *userService) Delete(id int64) error {
	err := dao.UserDao.UpdateColumn(simpleDb.DB(), id, "status", constants.StatusDeleted)
	if err == nil {
		// 删掉标签文章
		//UserTagService.DeleteByUserId(id)
	}
	return err
}

// 根据id批量获取
func (s *userService) GetUserInIds(UserIds []int64) []model.User {
	if len(UserIds) == 0 {
		return nil
	}
	var Users []model.User
	simpleDb.DB().Where("id in (?)", UserIds).Order("id desc").Find(&Users)
	return Users
}
