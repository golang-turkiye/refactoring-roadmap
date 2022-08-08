package shorturlhandler

import (
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/helpers/utils/authentication"
	responseutils "github.com/Golang-Turkiye/refactoring-roadmap/internal/helpers/utils/response"
	"github.com/Golang-Turkiye/refactoring-roadmap/internal/urlshorter/service"
	"github.com/Golang-Turkiye/refactoring-roadmap/src/response"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type shortURLHandler struct {
	linkService service.LinkService
	userService service.UserService
	router      *mux.Router
	logger      *logrus.Logger
}

func NewShortURLHander(linkService service.LinkService, userService service.UserService, router *mux.Router, logger *logrus.Logger) *shortURLHandler {
	return &shortURLHandler{
		linkService: linkService,
		userService: userService,
		router:      router,
		logger:      logger,
	}
}

// Run starts short url handler
func (h *shortURLHandler) Run() {
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
func (h *shortURLHandler) GoLink(w http.ResponseWriter, r *http.Request) {

}

// GetLink returns link by id
func (h *shortURLHandler) GetLink(w http.ResponseWriter, r *http.Request) {

}

// GetAllLinks returns all links
func (h *shortURLHandler) GetAllLinks(w http.ResponseWriter, r *http.Request) {
	response.OKResponse(w, "OK", h.logger)
}

// CreateLink creates link
func (h *shortURLHandler) CreateLink(w http.ResponseWriter, r *http.Request) {

}

// DeactivateLink deactivates link by id
func (h *shortURLHandler) DeactivateLink(w http.ResponseWriter, r *http.Request) {

}

// Login returns user by id
func (h *shortURLHandler) Login(w http.ResponseWriter, r *http.Request) {

}

// GetUser returns user by id
func (h *shortURLHandler) GetUser(w http.ResponseWriter, r *http.Request) {
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
	userResponse := responseutils.MapUserResponse(user)
	response.OKResponse(w, userResponse, h.logger)
}
func (h *shortURLHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {

}
