package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest-api-service/internal/handlers"
	"rest-api-service/pkg/logging"
)

var _ handlers.Handler = &handler{} // hint:

const (
	usersURL = "/api/users"
	userURL  = "/api/users/:id"
)

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{
		logger: logger,
	}
}

type handler struct {
	logger *logging.Logger
}

func (h *handler) Register(router *httprouter.Router) {
	h.logger.Info("Registering user handlers")
	router.GET("/api/users", h.GetUsers)
	router.GET("/api/users/:id", h.GetUser)
	router.POST("/api/users", h.CreateUser)
	router.PUT("/api/users/:id", h.UpdateUser)
	router.DELETE("/api/users/:id", h.DeleteUser)
	router.PATCH("/api/users/:id", h.PatchUser)
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("Users: "))
}

func (h *handler) GetUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello " + ps.ByName("id")))
}

func (h *handler) CreateUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Hello " + ps.ByName("id")))
}

func (h *handler) DeleteUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("Hello " + ps.ByName("id")))
}

func (h *handler) PatchUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello " + ps.ByName("id")))
}

func (h *handler) UpdateUser(w http.ResponseWriter, request *http.Request, ps httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hello " + ps.ByName("id")))
}
