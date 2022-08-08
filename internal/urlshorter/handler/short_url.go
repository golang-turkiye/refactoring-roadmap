package handler

import "net/http"

type ShortURLHandler interface {
	Run()
	GoLink(w http.ResponseWriter, r *http.Request)
	GetLink(w http.ResponseWriter, r *http.Request)
	GetAllLinks(w http.ResponseWriter, r *http.Request)
	CreateLink(w http.ResponseWriter, r *http.Request)
	DeactivateLink(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	GetUserByEmail(w http.ResponseWriter, r *http.Request)
}
