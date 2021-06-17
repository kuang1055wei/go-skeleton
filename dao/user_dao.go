package dao

import (
	"go-skeleton/model"
	"go-skeleton/pkg/simpleDb"

	"gorm.io/gorm"
)

var UserDao = newUserDao()

func newUserDao() *userDao {
	return &userDao{}
}

type userDao struct {
}

func (c *userDao) Get(db *gorm.DB, id int64) *model.User {
	code := &model.User{}
	res := db.First(code, id)
	if res.Error != nil {
		return nil
	}
	return code
}

func (c *userDao) Take(db *gorm.DB, where ...interface{}) *model.User {
	ret := &model.User{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userDao) Find(db *gorm.DB, cnd *simpleDb.SqlCnd) (list []model.User) {
	cnd.Find(db, &list)
	return
}

func (r *userDao) FindOne(db *gorm.DB, cnd *simpleDb.SqlCnd) *model.User {
	ret := &model.User{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (c *userDao) Create(db *gorm.DB, t *model.User) (err error) {
	err = db.Create(t).Error
	return
}

func (c *userDao) Update(db *gorm.DB, t *model.User) (err error) {
	err = db.Save(t).Error
	return
}

func (c *userDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (c *userDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.User{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (c *userDao) Delete(db *gorm.DB, id int64) {
	db.Delete(&model.User{}, "id = ?", id)
}

// BatchSave 批量插入数据
func (c *userDao) BatchSave(db *gorm.DB, courseCode []*model.User) error {
	return db.CreateInBatches(courseCode, len(courseCode)).Error
}

func (c *userDao) FindPageByParams(db *gorm.DB, params *simpleDb.QueryParams) (list []model.User, paging *simpleDb.Paging) {
	return c.FindPageByCnd(db, &params.SqlCnd)
}

func (c *userDao) FindPageByCnd(db *gorm.DB, cnd *simpleDb.SqlCnd) (list []model.User, paging *simpleDb.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.User{})

	paging = &simpleDb.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userDao) GetByUsername(db *gorm.DB, username string) *model.User {
	return r.Take(db, "username = ?", username)
}
