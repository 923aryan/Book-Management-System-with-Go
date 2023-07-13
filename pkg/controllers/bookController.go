package controllers

import (
	"bms/pkg/models"
	"bms/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	fmt.Println("accessed get book")
	res, err := json.Marshal(newBooks)

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)
	if err != nil {
		return
	}
	newBook, _ := models.GetBooksById(bookId)

	res, err := json.Marshal(newBook)

	if err != nil {
		fmt.Println("error while parsing")

	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	bookId, err := strconv.ParseInt(params["bookId"], 0, 0)

	if err != nil {
		fmt.Println("error while parsing")

	}

	newBook := models.DeleteBook(bookId)

	res, err := json.Marshal(newBook)
	if err != nil {
		fmt.Println("error while parsing")

	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	bookModel := &models.Book{}

	//The reason why passing bookModel without reference works is because bookModel is already a pointer to a Book struct type,
	// as you declared it with the & operator: bookModel := &models.Book{}
	utils.PareseBody(r, bookModel)
	b := bookModel.CreateBook()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.PareseBody(r, updateBook)
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBooksById(ID)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Auhtor != "" {
		bookDetails.Auhtor = updateBook.Auhtor
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
