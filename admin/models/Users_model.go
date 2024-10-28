package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func (user User) Migrate() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&user)
	fmt.Println("Database migrated")
}

func (user User) Add() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Veritabanı bağlantı hatası:", err)
		return nil
	}

	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println("Veritabanına ekleme hatası:", result.Error)
		return nil
	}
	return nil
}

func (user User) Get(where ...interface{}) User {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return user
	}
	db.First(&user, where...)
	return user
}

func (user User) GetAll(where ...interface{}) []User {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var posts []User
	db.Find(&posts, where...)
	return posts
}

func (user User) Update(column string, value interface{}) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&user).Update(column, value)
}

func (user User) Updates(data User) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&user).Updates(data)
}

func (user User) Delete() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&user, user.ID)
}
