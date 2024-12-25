package user

type User struct {
	ID        uint
	Email     string `gorm:"unique"`
	Password  string
	Name      string
	CreatedAt int64
}

func (User) TableName() string {
	return "users"
}
