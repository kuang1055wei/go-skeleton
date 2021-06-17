package services

import (
	"errors"
	"go-skeleton/dao"
	"go-skeleton/model"
	"go-skeleton/model/constants"
	"go-skeleton/pkg/simpleDb"
	"go-skeleton/utils"
	"strings"
)

var UserService = newUserService()

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

// GetByUsername 根据用户名查找
func (s *userService) GetByUsername(username string) *model.User {
	return dao.UserDao.GetByUsername(simpleDb.DB(), username)
}

// isUsernameExists 用户名是否存在
func (s *userService) isUsernameExists(username string) bool {
	return s.GetByUsername(username) != nil
}

//注册
func (s *userService) SignUp(username, password, rePassword string) (*model.User, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	err := utils.IsPassword(password, rePassword)
	if err != nil {
		return nil, err
	}

	if len(username) > 0 {
		if err := utils.IsUsername(username); err != nil {
			return nil, err
		}
		if s.isUsernameExists(username) {
			return nil, errors.New("用户名[" + username + "]已被占用")
		}
	}

	user := &model.User{
		Username: username,
		Password: utils.EncodePassword(password),
		Role:     1,
	}

	err = dao.UserDao.Create(simpleDb.DB(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
