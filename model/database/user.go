package database

type User struct {
	UserId   int    `gorm:"primaryKey;autoIncrement;not null;comment:用户ID" json:"user_id"`
	Username string `gorm:"type:varchar(30);not null;unique;comment:用户名" json:"username"`
	Password string `gorm:"type:varchar(30);not null;comment:密码" json:"password"`
	Nickname string `gorm:"type:varchar(30);not null;comment:昵称" json:"nickname"`
	UserType string `gorm:"type:varchar(30);not null;comment:用户类型" json:"user_type"`
	Email    string `gorm:"type:varchar(30);not null;comment:邮箱" json:"email"`
}
