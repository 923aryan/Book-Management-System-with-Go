// use -> constructing the book model
package models

import (
	"github.com/jinzhu/gorm"

	"bms/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Auhtor      string `json:"author"`
	Publication string `json:"Publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	// making the db to take the form of book struct(structure)
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBooksById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("Id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("Id=?", ID).Delete(&book)
	return book
}
