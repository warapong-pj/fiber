package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"name;unique"`
	Firstname string `gorm:"firstname"`
	Lastname  string `gorm:"lastname"`
	Email     string `gorm:"email;unique"`
}

type UserDomain interface {
	Reads()
	Read()
	Create()
	Update()
	Delete()
}

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) userRepositoryDB {
	return userRepositoryDB{db: db}
}
