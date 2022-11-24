package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-service/internal/handlers"
)

var _ handlers.Handler = &handler{} // hint:

const (
	usersUrl = "/api/users"
	userURL  = "/api/users/:id"
)

func NewHandler() handlers.Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Users: "))
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET("/api/users", h.GetUsers)
	router.GET("/api/users/:id", h.GetUser)
	router.POST("/api/users", h.CreateUser)
	router.PUT("/api/users/:id", h.UpdateUser)
	router.DELETE("/api/users/:id", h.DeleteUser)
	router.PATCH("/api/users/:id", h.PatchUser)
}

func (h *handler) GetUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello " + ps.ByName("id")))
	w.WriteHeader(200)
}

func (h *handler) CreateUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello " + ps.ByName("id")))
	w.WriteHeader(204)
}

func (h *handler) DeleteUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello " + ps.ByName("id")))
	w.WriteHeader(204)
}

func (h *handler) PatchUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello " + ps.ByName("id")))
	w.WriteHeader(200)
}

func (h *handler) UpdateUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello " + ps.ByName("id")))
	w.WriteHeader(200)
}
