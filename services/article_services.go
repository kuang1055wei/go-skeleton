package services

import (
	"go-skeleton/dao"
	"go-skeleton/model"
	"go-skeleton/model/constants"
	"go-skeleton/pkg/simpleDb"

	"gorm.io/gorm"
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

func (s *articleService) CreateArticle(article *model.Article) error {
	return dao.ArticleDao.Create(simpleDb.DB(), article)
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

//以下方法没有优化，都是刚学习gorm时候的demo
//使用gorm.Expr使用表达式
func (a *articleService) IncrReadCount(id int) error {
	var art model.Article
	return simpleDb.DB().Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1)).Error
}

//下面这些函数 都可以迁移到services中去
// GetFromID 通过id获取内容 Primary key
func (a *articleService) GetArticleById(id int) (model.Article, error) {
	var result model.Article
	//First、Take、Last  没有找到记录时，它会返回 ErrRecordNotFound 错误
	res := simpleDb.DB().Preload("Category").First(&result, id)

	//res := db.Table(obj.TableName()).Preload("Category").Where("id = ?", id).Find(&result)
	//res.RowsAffected //返回找到的记录数，相当于 `len(users)`
	if res.Error != nil {
		return result, res.Error
	}
	return result, res.Error
}

//获取文章列表
func (a *articleService) GetArticleList(condition map[string]string, pageSize int, page int) ([]model.Article, int64) {
	var articleList []model.Article
	var total int64
	column := "id,title,`cid`,`desc`,img,comment_count,read_count,created_at,updated_at"
	//column := "id,title, img,  `desc`, comment_count, read_count"
	tx := simpleDb.DB().Table("article").
		Preload("Category").
		Select(column).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("created_at desc")

	if _, ok := condition["title"]; ok {
		tx.Where("title like ?", condition["title"]+"%")
	}

	err := tx.Find(&articleList).Count(&total).Error
	if err != nil {
		return nil, 0
	}

	return articleList, total
}

func (a *articleService) GetArticleList2(condition map[string]string, pageSize int, page int) ([]map[string]interface{}, int64) {
	var articleList []map[string]interface{}
	var total int64
	column := "id,title,created_at"
	//column := "id,title, img,  `desc`, comment_count, read_count"
	tx := simpleDb.DB().Table("article").
		Select(column).
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("created_at desc")

	if _, ok := condition["title"]; ok {
		tx.Where("title like ?", condition["title"]+"%")
	}

	err := tx.Find(&articleList).Count(&total).Error
	if err != nil {
		return nil, 0
	}

	return articleList, total
}

func (a *articleService) SearchArticle(title string, pageSize int, page int) ([]model.Article, int64) {
	var articleList []model.Article
	var total int64
	column := "id,title,cid,`desc`,img,comment_count,read_count,created_at,updated_at"
	err := simpleDb.DB().Table("article").
		Select(column).
		Where("title Like ?", title+"%").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("created_at desc").
		Find(&articleList).
		Count(&total).Error
	if err != nil {
		return nil, 0
	}
	return articleList, total
}

func (a *articleService) SearchArticle2(title string, pageSize int, page int) ([]model.Article, int64, error) {
	var articleList []model.Article
	var total int64
	column := "id,title,cid,`desc`,img,comment_count,read_count,created_at,updated_at"
	result := simpleDb.DB().Table("article").
		Select(column).
		Where("title Like ?", title+"%").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Order("created_at desc").
		Find(&articleList).
		Count(&total)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return articleList, total, result.Error
}

func (a *articleService) createArticle(data *model.Article) (bool, error) {
	err := simpleDb.DB().Create(&data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *articleService) EditArticle(id int, data *model.Article) (bool, error) {
	var art model.Article
	var column = make(map[string]interface{})
	column["title"] = data.Title
	column["cid"] = data.Cid
	column["desc"] = data.Desc
	column["content"] = data.Content
	column["img"] = data.Img
	result := simpleDb.DB().Debug().Model(&art).Where("id=?", id).Updates(column)
	if result.Error != nil {
		return false, result.Error
	}
	//if result.Error.(ErrRecordNotFound) {
	//
	//}
	return true, nil
}

// 删除文章
func (a *articleService) DeleteArt(id int) (bool, error) {
	var art model.Article
	err := simpleDb.DB().Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
