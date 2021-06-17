package model

import "gorm.io/gorm"

// User [...]
type User struct {
	gorm.Model
	Username string `gorm:"column:username;type:varchar(20);not null" json:"username"`
	Password string `gorm:"column:password;type:varchar(500);not null" json:"-"`
	Role     *int64 `gorm:"column:role;type:bigint(20);default:2" json:"role"`
}

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID        string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
	Username  string
	Password  string
	Role      string
}{
	ID:        "id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
	Username:  "username",
	Password:  "password",
	Role:      "role",
}
