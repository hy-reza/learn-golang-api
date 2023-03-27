package service

import (
	"books-api/model"
	"books-api/repository"
	"books-api/types"
	"log"
)

type Service interface {
	FindAll() ([]model.Book, error)
	FindByID(ID int) (model.Book, error)
	Create(b types.BookRequest) (model.Book, error)
	Update(ID int, b types.BookRequest) (model.Book, error)
	Delete(ID int) (model.Book, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{repo}
}

func (s *service) FindAll() ([]model.Book, error) {
	books, err := s.repo.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (model.Book, error) {
	book, err := s.repo.FindByID(ID)
	return book, err
}

func (s *service) Create(book types.BookRequest) (model.Book, error) {
	b := model.Book{
		Author: book.Author,
		Title:  book.Title,
		Desc:   book.Desc,
	}

	newBook, err := s.repo.Create(b)

	return newBook, err
}

func (s *service) Update(ID int, bookReq types.BookRequest) (model.Book, error) {
	book, err := s.repo.FindByID(ID)
	if err != nil {
		log.Fatal(err)
	}

	book.Author = bookReq.Author
	book.Title = bookReq.Title
	book.Desc = bookReq.Desc

	newBook, err := s.repo.Update(book)

	return newBook, err
}

func (s *service) Delete(ID int) (model.Book, error) {
	book, err := s.repo.FindByID(ID)
	if err != nil {
		log.Fatal(err)
	}

	newBook, err := s.repo.Delete(book)

	return newBook, err
}
