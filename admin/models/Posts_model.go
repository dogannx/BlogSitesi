package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Picture_url string
	Category_id                                    int
}

func (post Post) Migrate() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&post)
	fmt.Println("Database migrated")
}

func (post Post) Add() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Veritabanı bağlantı hatası:", err)
		return nil
	}

	result := db.Create(&post)
	if result.Error != nil {
		fmt.Println("Veritabanına ekleme hatası:", result.Error)
		return nil
	}
	return nil
}

func (post Post) Get(where ...interface{}) Post {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return post
	}
	db.First(&post, where...)
	return post
}

func (post Post) GetAll(where ...interface{}) []Post {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var posts []Post
	db.Find(&posts, where...)
	return posts
}

func (post Post) Update(column string, value interface{}) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Update(column, value)
}

func (post Post) Updates(data Post) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&post).Updates(data)
}

func (post Post) Delete() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&post, post.ID)
}
