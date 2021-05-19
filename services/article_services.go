package services

import (
	"gin-test/dao"
	"gin-test/model"
	"gin-test/model/constants"
	"gin-test/pkg/simpleDb"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

func (s *articleService) Get(id int64) *model.Article {
	return dao.ArticleDao.Get(simpleDb.DB(), id)
}

func (s *articleService) Take(where ...interface{}) *model.Article {
	return dao.ArticleDao.Take(simpleDb.DB(), where...)
}

func (s *articleService) Find(cnd *simpleDb.SqlCnd) []model.Article {
	return dao.ArticleDao.Find(simpleDb.DB(), cnd)
}

func (s *articleService) FindOne(cnd *simpleDb.SqlCnd) *model.Article {
	return dao.ArticleDao.FindOne(simpleDb.DB(), cnd)
}

func (s *articleService) FindPageByParams(params *simpleDb.QueryParams) (list []model.Article, paging *simpleDb.Paging) {
	return dao.ArticleDao.FindPageByParams(simpleDb.DB(), params)
}

func (s *articleService) FindPageByCnd(cnd *simpleDb.SqlCnd) (list []model.Article, paging *simpleDb.Paging) {
	return dao.ArticleDao.FindPageByCnd(simpleDb.DB(), cnd)
}

func (s *articleService) Update(t *model.Article) error {
	err := dao.ArticleDao.Update(simpleDb.DB(), t)
	return err
}

func (s *articleService) Updates(id int64, columns map[string]interface{}) error {
	err := dao.ArticleDao.Updates(simpleDb.DB(), id, columns)
	return err
}

func (s *articleService) UpdateColumn(id int64, name string, value interface{}) error {
	err := dao.ArticleDao.UpdateColumn(simpleDb.DB(), id, name, value)
	return err
}

func (s *articleService) Delete(id int64) error {
	err := dao.ArticleDao.UpdateColumn(simpleDb.DB(), id, "status", constants.StatusDeleted)
	if err == nil {
		// 删掉标签文章
		//ArticleTagService.DeleteByArticleId(id)
	}
	return err
}

// 根据文章编号批量获取文章
func (s *articleService) GetArticleInIds(articleIds []int64) []model.Article {
	if len(articleIds) == 0 {
		return nil
	}
	var articles []model.Article
	simpleDb.DB().Where("id in (?)", articleIds).Order("id desc").Find(&articles)
	return articles
}
