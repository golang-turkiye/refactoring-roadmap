package apiV1

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/service"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/core/usecase"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/authentication"
	"github.com/Golang-Turkiye/refactoring-roadmap/pkg/response"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ShortURLHandler struct {
	linkService service.LinkService
	userService service.UserService
	router      *mux.Router
	logger      *logrus.Logger
}

func NewShortURLHander(linkService service.LinkService, userService service.UserService, router *mux.Router, logger *logrus.Logger) *ShortURLHandler {
	return &ShortURLHandler{
		linkService: linkService,
		userService: userService,
		router:      router,
		logger:      logger,
	}
}

// Run starts short url handler
func (h *ShortURLHandler) Run() {
	h.router.HandleFunc("/v1/link/all", h.GetAllLinks).Methods("GET")
	h.router.HandleFunc("/v1/link/", h.CreateLink).Methods("POST")
	h.router.HandleFunc("/v1/link/{shortPath}", h.GetLink).Methods("GET")
	h.router.HandleFunc("/v1/link/{shortPath}/deactivate", h.GetLink).Methods("GET")
	h.router.HandleFunc("/v1/{shortPath}", h.GoLink).Methods("GET")
	h.router.HandleFunc("/v1/user/login", h.Login).Methods("POST")
	h.router.HandleFunc("/v1/user/{userID}", h.GetUser).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: h.router,
	}
	h.logger.Info("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil {
		h.logger.WithFields(logrus.Fields{
			"error": err,
		}).Error("Error starting server")
		return
	}
}

// GoLink redirects to link
func (h *ShortURLHandler) GoLink(w http.ResponseWriter, r *http.Request) {

}

// GetLink returns link by id
func (h *ShortURLHandler) GetLink(w http.ResponseWriter, r *http.Request) {

}

// GetAllLinks returns all links
func (h *ShortURLHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	response.OKResponse(w, "OK", h.logger)
}

// CreateLink creates link
func (h *ShortURLHandler) CreateLink(w http.ResponseWriter, r *http.Request) {

}

// DeactivateLink deactivates link by id
func (h *ShortURLHandler) DeactivateLink(w http.ResponseWriter, r *http.Request) {

}

// Login returns user by id
func (h *ShortURLHandler) Login(w http.ResponseWriter, r *http.Request) {

}

// GetUser returns user by id
func (h *ShortURLHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		h.logger.Error("Token is empty on GetUser")
		response.UnauthorizedResponse(w, "Error getting user", h.logger)
		return
	}
	email, err := authentication.GetEmailByToken(token)
	if err != nil {
		h.logger.Error("Authorization error on getting user")
		response.UnauthorizedResponse(w, "Error getting user", h.logger)
		return
	}
	user, err := h.userService.GetUserByEmail(email)
	if err != nil {
		response.UnauthorizedResponse(w, "Not Authorized", h.logger)
		return
	}
	userResponse := usecase.MapUserResponse(user)
	response.OKResponse(w, userResponse, h.logger)
}
func (h *ShortURLHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

}
