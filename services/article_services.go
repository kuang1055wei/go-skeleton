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

// 获取文章对应的标签
//func (s *articleService) GetArticleTags(articleId int64) []model.Tag {
//	articleTags := dao.ArticleDao.Find(simpleDb.DB(), simpleDb.NewSqlCnd().Where("article_id = ?", articleId))
//	var tagIds []int64
//	for _, articleTag := range articleTags {
//		tagIds = append(tagIds, articleTag.TagId)
//	}
//	return cache.TagCache.GetList(tagIds)
//}

// 文章列表
//func (s *articleService) GetArticles(cursor int64) (articles []model.Article, nextCursor int64) {
//	cnd := simpleDb.NewSqlCnd().Eq("status", constants.StatusOk).Desc("id").Limit(20)
//	if cursor > 0 {
//		cnd.Lt("id", cursor)
//	}
//	articles = dao.ArticleDao.Find(simpleDb.DB(), cnd)
//	if len(articles) > 0 {
//		nextCursor = articles[len(articles)-1].Id
//	} else {
//		nextCursor = cursor
//	}
//	return
//}

// 标签文章列表
func (s *articleService) GetTagArticles(tagId int64, cursor int64) (articles []model.Article, nextCursor int64) {
	cnd := simpleDb.NewSqlCnd().Eq("tag_id", tagId).Eq("status", constants.StatusOk).Desc("id").Limit(20)
	if cursor > 0 {
		cnd.Lt("id", cursor)
	}
	nextCursor = cursor
	articleTags := dao.ArticleDao.Find(simpleDb.DB(), cnd)
	if len(articleTags) > 0 {
		var articleIds []int64
		//for _, articleTag := range articleTags {
		//	articleIds = append(articleIds, articleTag.ArticleId)
		//	nextCursor = articleTag.Id
		//}
		articles = s.GetArticleInIds(articleIds)
	}
	return
}

// 发布文章
//func (s *articleService) Publish(userId int64, title, summary, content, contentType string, tags []string,
//	sourceUrl string) (article *model.Article, err error) {
//	title = strings.TrimSpace(title)
//	summary = strings.TrimSpace(summary)
//	content = strings.TrimSpace(content)
//
//	if simpleDb.IsBlank(title) {
//		return nil, errors.New("标题不能为空")
//	}
//	if simpleDb.IsBlank(content) {
//		return nil, errors.New("内容不能为空")
//	}
//
//	// 获取后台配置 否是开启发表文章审核
//	status := constants.StatusOk
//	sysConfigArticlePending := cache.SysConfigCache.GetValue(constants.SysConfigArticlePending)
//	if strings.ToLower(sysConfigArticlePending) == "true" {
//		status = constants.StatusPending
//	}
//
//	article = &model.Article{
//		UserId:      userId,
//		Title:       title,
//		Summary:     summary,
//		Content:     content,
//		ContentType: contentType,
//		Status:      status,
//		SourceUrl:   sourceUrl,
//		CreateTime:  date.NowTimestamp(),
//		UpdateTime:  date.NowTimestamp(),
//	}
//
//	err = simpleDb.DB().Transaction(func(tx *gorm.DB) error {
//		tagIds := repositories.TagRepository.GetOrCreates(tx, tags)
//		err := dao.ArticleDao.Create(tx, article)
//		if err != nil {
//			return err
//		}
//		dao.ArticleDao.AddArticleTags(tx, article.Id, tagIds)
//		return nil
//	})
//
//	if err == nil {
//		seo.Push(urls.ArticleUrl(article.Id))
//	}
//	return
//}

// 修改文章
//func (s *articleService) Edit(articleId int64, tags []string, title, content string) *simpleDb.CodeError {
//	if len(title) == 0 {
//		return simpleDb.NewErrorMsg("请输入标题")
//	}
//	if len(content) == 0 {
//		return simpleDb.NewErrorMsg("请填写文章内容")
//	}
//
//	err := simpleDb.DB().Transaction(func(tx *gorm.DB) error {
//		err := dao.ArticleDao.Updates(simpleDb.DB(), articleId, map[string]interface{}{
//			"title":   title,
//			"content": content,
//		})
//		if err != nil {
//			return err
//		}
//		tagIds := repositories.TagRepository.GetOrCreates(tx, tags) // 创建文章对应标签
//		dao.ArticleDao.DeleteArticleTags(tx, articleId)             // 先删掉所有的标签
//		dao.ArticleDao.AddArticleTags(tx, articleId, tagIds)        // 然后重新添加标签
//		return nil
//	})
//	cache.ArticleTagCache.Invalidate(articleId)
//	return simpleDb.FromError(err)
//}