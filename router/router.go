package router

import (
	"github.com/gorilla/mux"
	"github.com/pavanilla/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	//User collections
	// create new user
	router.HandleFunc("/users/user", middleware.CreateUser).Methods("POST")
	//EDit user DEtails
	//  the id should be a valid id
	router.HandleFunc("/users?id={12345}", middleware.EditUser).Methods("PUT")
	//Get users
	router.HandleFunc("/users/getUser", middleware.GetUser).Methods("GET")

	//search Users
	//  the id should be a valid id
	router.HandleFunc("/users?id={12345}", middleware.SearchUser).Methods("GET")
	//Login endpoint
	router.HandleFunc("/users?name={pavan}&password={pavanilla}", middleware.LoginEndPoint).Methods("GET")

	//posts  collections
	// create posts
	router.HandleFunc("/posts/post", middleware.CreatePost).Methods("POST")
	//edit posts
	router.HandleFunc("/posts/edit?id={}", middleware.EditPost).Methods("PUT")
	//List all posts
	router.HandleFunc("/posts/list", middleware.Allposts).Methods("GET", "options")
	//Delete posts
	router.HandleFunc("/posts/delete?id={}", middleware.DeletePosts).Methods("DELETE")
	return router

}
