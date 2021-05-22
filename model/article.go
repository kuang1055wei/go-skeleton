package model

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
