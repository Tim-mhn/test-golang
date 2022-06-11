package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tim-mhn/module-1/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)

		case http.MethodPost:
			uc.add(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}

	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)

		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
		}

		id, err := strconv.Atoi(matches[1])

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(w, r, id)
		case http.MethodDelete:
			uc.delete(w, id)

		default:
			w.WriteHeader(http.StatusNotFound)

		}

	}
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	users := models.GetUsers()
	EncodeResponseAsJson(users, w)
}

func (uc userController) get(w http.ResponseWriter, r *http.Request, id int) {
	user, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	EncodeResponseAsJson(user, w)
}

func (uc userController) add(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(w, r)
	if err == nil {
		models.AddUser(u)
	}
}

func (uc userController) delete(w http.ResponseWriter, id int) {
	err := models.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

	}
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)`),
	}
}

func (uc *userController) parseRequest(w http.ResponseWriter, r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}

	return u, nil
}
