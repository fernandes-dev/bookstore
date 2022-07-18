package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fernandes-dev/go-bookstore/pkg/models"
	"github.com/fernandes-dev/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ResponseMessage struct {
	Message string `json:"message"`
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	
	if db.RowsAffected == 0 {
		errorMessage := ResponseMessage{Message: "Book not found"}
		jsonResponse, _ := json.Marshal(errorMessage)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResponse)
		return
	}

	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	_, db := models.DeleteBook(ID)

	if db.RowsAffected == 0 {
		errorMessage := ResponseMessage{Message: "Book not found"}
		jsonResponse, _ := json.Marshal(errorMessage)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResponse)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)

	if db.RowsAffected == 0 {
		errorMessage := ResponseMessage{Message: "Book not found"}
		jsonResponse, _ := json.Marshal(errorMessage)
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonResponse)
		return
	}

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
