package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (category Category) Migrate() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&category)
	fmt.Println("Database migrated")
}

func (category Category) Add() error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Veritabanı bağlantı hatası:", err)
		return nil
	}

	result := db.Create(&category)
	if result.Error != nil {
		fmt.Println("Veritabanına ekleme hatası:", result.Error)
		return nil
	}
	return nil
}

func (category Category) Get(where ...interface{}) Category {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return category
	}
	db.First(&category, where...)
	return category
}

func (category Category) GetAll(where ...interface{}) []Category {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var categories []Category
	db.Find(&categories, where...)
	return categories

}

func (category Category) Update(column string, value interface{}) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&category).Update(column, value)
}

func (category Category) Updates(data Category) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Model(&category).Updates(data)
}

func (category Category) Delete() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&category, category.ID)
}
