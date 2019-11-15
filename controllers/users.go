package controllers

import (
	"api/database"
	"api/models"
	responses "api/reponses"
	"api/repository"
	"api/repository/crud"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)
	func (usersRepository repository.UserRepository) {
		users, err := usersRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusOK, users)
	}(repo)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryUsersCRUD(db)
	func (usersRepository repository.UserRepository) {
		user, err = usersRepository.Save(user)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := mux.Vars(r)["id"]
	newUserId , newUserIdFormatError := strconv.ParseUint(string(userId), 10, 32)
	if err != true {
		responses.ERROR(w, http.StatusBadRequest, errors.New("user id not supplied in the request"))
		return
	}
	if newUserIdFormatError != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("parameter supplied is of invalid format"))
		return
	}
	//responses.JSON(w, http.StatusCreated, userId)
	db, dbError := database.Connect()
	if dbError != nil {
		responses.ERROR(w, http.StatusInternalServerError, dbError)
		return
	}
	repo := crud.NewRepositoryUsersCRUD(db)
	func (usersRepository repository.UserRepository) {
		user, repoError := usersRepository.FindById(uint32(newUserId))
		if repoError != nil {
			responses.ERROR(w, http.StatusNotFound, repoError)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, user.ID))
		responses.JSON(w, http.StatusCreated, user)
	}(repo)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := mux.Vars(r)["id"]
	newUserId , newUserIdFormatError := strconv.ParseUint(string(userId), 10, 32)
	if err != true {
		responses.ERROR(w, http.StatusBadRequest, errors.New("user id not supplied in the request"))
		return
	}
	if newUserIdFormatError != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("parameter supplied is of invalid format"))
		return
	}
	//responses.JSON(w, http.StatusCreated, userId)
	db, dbError := database.Connect()
	if dbError != nil {
		responses.ERROR(w, http.StatusInternalServerError, dbError)
		return
	}
	repo := crud.NewRepositoryUsersCRUD(db)
	func (usersRepository repository.UserRepository) {
		rows, repoError := usersRepository.Update(uint32(newUserId))
		if repoError != nil {
			responses.ERROR(w, http.StatusNotFound, repoError)
			return
		}

		if rows >= 1 {
			w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
			responses.JSON(w, http.StatusOK, map[string]interface{}{
				"Message": "User record successfully updated",
			})
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"Message": "User record could not be updated.",
		})
		return
	}(repo)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := mux.Vars(r)["id"]
	newUserId , newUserIdFormatError := strconv.ParseUint(string(userId), 10, 32)
	if err != true {
		responses.ERROR(w, http.StatusBadRequest, errors.New("user id not supplied in the request"))
		return
	}
	if newUserIdFormatError != nil {
		responses.ERROR(w, http.StatusBadRequest, errors.New("parameter supplied is of invalid format"))
		return
	}
	db, dbError := database.Connect()
	if dbError != nil {
		responses.ERROR(w, http.StatusInternalServerError, dbError)
		return
	}
	repo := crud.NewRepositoryUsersCRUD(db)
	func (usersRepository repository.UserRepository) {
		rows, repoError := usersRepository.Delete(uint32(newUserId))
		if repoError != nil {
			responses.ERROR(w, http.StatusNotFound, repoError)
			return
		}

		if rows >= 1 {
			w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
			responses.JSON(w, http.StatusOK, map[string]interface{}{
				"Message": "User record successfully deleted",
			})
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusInternalServerError, map[string]interface{}{
			"Message": "User record could not be modified.",
		})
		return
	}(repo)
}
