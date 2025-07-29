package db

import (
	"github.com/MarceloBxD/bookstore/internal/domain/book"
	"gorm.io/gorm"
)

type GormBookRepository struct {
	db *gorm.DB 
}

func NewGormBookRepository(db *gorm.DB) *GormBookRepository {
	return &GormBookRepository{db: db}
}

func (r *GormBookRepository) Create(b *book.Book) error {
	return r.db.Create(b).Error 
}

func (r *GormBookRepository) FindById(id uint) (*book.Book, error) {
	var b book.Book 
	err := r.db.First(&b, id).Error 
	return &b, err 
}

func (r *GormBookRepository) FindAll() ([]book.Book, error) {
	var books []book.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *GormBookRepository) Update(b *book.Book) error {
	return r.db.Save(b).Error
}

func (r *GormBookRepository) Delete(id uint) error {
	return r.db.Delete(&book.Book{}, id).Error
}