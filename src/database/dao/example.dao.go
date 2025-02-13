package dao

// Entity

type UserEntity struct {
	ID       *uint     `gorm:"primarykey"`
	Name     string    `gorm:"size:255;not null;"`
	Email    string    `gorm:"size:255;not null;uniqueIndex"`
}

func (UserEntity) TableName() string {
	return "users"
}

