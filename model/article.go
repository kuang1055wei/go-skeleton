package model

import (
	"gin-test/pkg/simpleDb"

	"gorm.io/gorm"
)

type Article struct {
	Model
	Title        string   `gorm:"column:title;type:varchar(100);not null" json:"title"`
	Cid          uint64   `gorm:"index:fk_article_category;column:cid;type:bigint(20) unsigned;not null" json:"cid"`
	Desc         string   `gorm:"column:desc;type:varchar(200)" json:"desc"`
	Content      string   `gorm:"column:content;type:longtext" json:"content"`
	Img          string   `gorm:"column:img;type:varchar(100)" json:"img"`
	CommentCount int64    `gorm:"column:comment_count;type:bigint(20);not null;default:0" json:"comment_count"`
	ReadCount    int64    `gorm:"column:read_count;type:bigint(20);not null;default:0" json:"read_count"`
	Category     Category `gorm:"foreignKey:Cid" json:"category"`
}

// ArticleColumns get sql column name.获取数据库列名
var ArticleColumns = struct {
	ID           string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	Title        string
	Cid          string
	Desc         string
	Content      string
	Img          string
	CommentCount string
	ReadCount    string
}{
	ID:           "id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	Title:        "title",
	Cid:          "cid",
	Desc:         "desc",
	Content:      "content",
	Img:          "img",
	CommentCount: "comment_count",
	ReadCount:    "read_count",
}

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "article"
}

//下面这些函数 都可以迁移到services中去
// GetFromID 通过id获取内容 Primary key
func (obj *Article) GetArticleById(id int) (Article, error) {
	var result Article
	//First、Take、Last  没有找到记录时，它会返回 ErrRecordNotFound 错误
	res := simpleDb.DB().Preload("Category").First(&result, id)

	//res := db.Table(obj.TableName()).Preload("Category").Where("id = ?", id).Find(&result)
	//res.RowsAffected //返回找到的记录数，相当于 `len(users)`
	if res.Error != nil {
		return result, res.Error
	}
	return result, res.Error
}

//使用gorm.Expr使用表达式
func IncrReadCount(id int) error {
	var art Article
	return simpleDb.DB().Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1)).Error
}

//获取文章列表
func GetArticleList(condition map[string]string, pageSize int, page int) ([]Article, int64) {
	var articleList []Article
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

func GetArticleList2(condition map[string]string, pageSize int, page int) ([]map[string]interface{}, int64) {
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

func SearchArticle(title string, pageSize int, page int) ([]Article, int64) {
	var articleList []Article
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

func SearchArticle2(title string, pageSize int, page int) ([]Article, int64, error) {
	var articleList []Article
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

func createArticle(data *Article) (bool, error) {
	err := simpleDb.DB().Create(&data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func EditArticle(id int, data *Article) (bool, error) {
	var art Article
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
func DeleteArt(id int) (bool, error) {
	var art Article
	err := simpleDb.DB().Where("id = ? ", id).Delete(&art).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
