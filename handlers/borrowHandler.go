package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	domain "library-sample/domains"
	"library-sample/router"
	"library-sample/services"
)

var borrowService = services.NewBorrowService(userService)

func listBorrows(w http.ResponseWriter, r *http.Request) {
	borrows := borrowService.ListBorrows()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrows)
}

func getBorrowByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrows/get/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid borrow ID", http.StatusBadRequest)
		return
	}

	borrow, err := borrowService.GetBorrowByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(borrow)
}

func addBorrow(w http.ResponseWriter, r *http.Request) {
	var borrow domain.Borrow
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &borrow)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	createdBorrow, err := borrowService.AddBorrow(borrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdBorrow)
}

func updateBorrow(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrow/update/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid borrow ID", http.StatusBadRequest)
		return
	}

	var updatedBorrow domain.Borrow
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &updatedBorrow)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	updatedBorrow.ID = id
	err = borrowService.UpdateBorrow(updatedBorrow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteBorrow(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/borrow/delete/")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid borrow ID", http.StatusBadRequest)
		return
	}

	err = borrowService.DeleteBorrow(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func BorrowRoutes() []router.Route {
	return []router.Route{
		{Path: "/borrow/list", Handler: listBorrows},
		{Path: "/borrow/get/", Handler: getBorrowByID},
		{Path: "/borrow/add", Handler: addBorrow},
		{Path: "/borrow/update/", Handler: updateBorrow},
		{Path: "/borrow/delete/", Handler: deleteBorrow},
	}
}
