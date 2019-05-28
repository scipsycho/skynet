package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	root "skynet/pkg"

	"github.com/gorilla/mux"
)

type userRouter struct {
	userService root.UserService
}

// NewUserRouter create the router for User schema
func NewUserRouter(u root.UserService, router *mux.Router) *mux.Router {
	userRouter := userRouter{u}

	router.HandleFunc("/create", userRouter.createUserHandler).Methods("POST")
	router.HandleFunc("/verify", userRouter.verifyUserHandler).Methods("POST")
	router.HandleFunc("/{username}", userRouter.getUserIdentifier).Methods("GET")

	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := root.User{UserName: r.FormValue("name"), Password: r.FormValue("password")}

	err := ur.userService.CreateUser(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, "/record/create", 302)
}

func (ur *userRouter) verifyUserHandler(w http.ResponseWriter, r *http.Request) {

	cred := root.Credentials{UserName: r.FormValue("name"), Password: r.FormValue("password")}
	res, err, flag := ur.userService.Login(cred)

	fmt.Println(res, flag)

	if !flag {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, "/display", 302)
}

func (ur *userRouter) getUserIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	username := vars["username"]

	user, err := ur.userService.GetUserByUsername(username)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

/*
// decodeUser parses the body of request
func decodeUser(r *http.Request) (root.User, error) {
	var u root.User

	if r.Body == nil {
		return u, errors.New("no request body")
	}

	err := json.NewDecoder(r.Body).Decode(&u)
	return u, err
}

*/
