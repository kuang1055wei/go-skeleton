package model

// Category [...]
type Category struct {
	ID   uint64 `gorm:"primaryKey" json:"-"`
	Name string `json:"name"`
}

// TableName get sql table name.获取数据库表名
func (m *Category) TableName() string {
	return "category"
}

// CategoryColumns get sql column name.获取数据库列名
var CategoryColumns = struct {
	ID   string
	Name string
}{
	ID:   "id",
	Name: "name",
}
