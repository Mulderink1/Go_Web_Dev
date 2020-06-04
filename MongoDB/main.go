package main

import (
	"github.com/Go_Web_Dev/MongoDB/controllers"
	"github.com/Go_Web_Dev/MongoDB/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8081", r)
}

func getSession() map[string]models.User {
	return models.LoadUsers()
}