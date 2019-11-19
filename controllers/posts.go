package controllers

import (
	"api/auth"
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

func CreatePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	fmt.Println("CHIKE ID", uid)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostsRepository) {
		post.Prepare()
		if err := post.Validate(); err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		post, err = postsRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostsRepository) {
		posts, err = postsRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusCreated, posts)
	}(repo)
}

func GetPostByID(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostsRepository) {
		post, err = postsRepository.FindById(pid)
		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	db, err := database.Connect()

	post := models.Post{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostsRepository) {
		_, err = postsRepository.Update(pid, post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusOK, map[string]string{
			"message": "Post updated",
		})
	}(repo)
}


func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	db, err := database.Connect()

	var uid uint32
	uid, err = auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
		return
	}

	post := models.Post{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	repo := crud.NewRepositoryPostsCRUD(db)
	func (postsRepository repository.PostsRepository) {
		_, err = postsRepository.DeletePost(pid, uid)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/", r.Host, r.RequestURI))
		responses.JSON(w, http.StatusOK, map[string]string{
			"message": "Post removed",
		})
	}(repo)
}