package repository

import (
	"books-api/model"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Book, error)
	FindByID(ID int) (model.Book, error)
	Create(b model.Book) (model.Book, error)
	Update(b model.Book) (model.Book, error)
	Delete(b model.Book) (model.Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]model.Book, error) {
	var books []model.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindByID(ID int) (model.Book, error) {
	var book model.Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) Create(book model.Book) (model.Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book model.Book) (model.Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}
func (r *repository) Delete(book model.Book) (model.Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
