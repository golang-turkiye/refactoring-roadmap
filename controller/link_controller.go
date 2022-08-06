package controller

import (
	"errors"
	"fmt"
	"github.com/Golang-Turkiye/refactoring-roadmap/db/mysql"
	"github.com/Golang-Turkiye/refactoring-roadmap/model"
	"math/rand"
	"net/http"
	"time"
)

func PostCreateLink(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDByToken(r)
	if err != nil {
		notAuthorizedResponse(w, err)
		return
	}
	conn, err := mysql.Connection()
	if err != nil {
		internalServerError(w, err)
		return
	}
	longURL := r.FormValue("longUrl")
	if longURL == "" {
		badRequestResponse(w, errors.New("longUrl is empty"))
		return
	}
	link := model.Link{
		OwnerID:        uint(userID),
		LongUrl:        longURL,
		ShortenURLPath: generateShortURLPath(),
		IsDeleted:      false,
	}
	if err := link.Save(conn); err != nil {
		internalServerError(w, err)
		return
	}
	successfulResponse(w, link)
	return
}

func generateShortURLPath() string {
	randomString := generateRandomString(6)
	return randomString
}

func generateRandomString(i int) string {
	randomString := ""
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for j := 0; j < i; j++ {
		randomString += string(rnd.Intn(26) + 65)
	}
	return randomString
}
func PostAllLinks(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDByToken(r)
	if err != nil {
		internalServerError(w, err)
		return
	}
	conn, err := mysql.Connection()
	if err != nil {
		internalServerError(w, err)
		return
	}
	links := model.Links{}
	links.FetchAll(conn, userID)

}
func GetGoLink(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDByToken(r)
	if err != nil {
		internalServerError(w, err)
		return
	}
	conn, err := mysql.Connection()
	if err != nil {
		internalServerError(w, err)
		return
	}
	shortPath := r.FormValue("path")
	link := model.Link{}
	link.FetchByCondition(conn, []string{fmt.Sprintf("user_id:%d", userID), fmt.Sprintf("shorten_url_path:%s", shortPath)})
	successfulResponse(w, link)
}
