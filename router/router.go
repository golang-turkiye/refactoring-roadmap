package router

import (
	"errors"
	"fmt"
	"github.com/Golang-Turkiye/refactoring-roadmap/controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RunHandlers() {
	router := mux.NewRouter()
	routers(router)
	log.Println(":8000 Port for Web Server")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatalln(err)
	}
}

func routers(router *mux.Router) error {
	err := errors.New("routers handler error")
	if e := userRouters(router); err != nil {
		err = fmt.Errorf("%w; %w", err, e)
	}
	if e := pathRouters(router); err != nil {
		err = fmt.Errorf("%w; %w", err, e)
	}
	return err
}

func pathRouters(router *mux.Router) error {
	pathRouter, err := getSubrouter(router, "link")
	if err != nil {
		return err
	}
	pathRouter.HandleFunc("", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("12313")
	}).Methods("GET")
	pathRouter.HandleFunc("", controller.PostCreateLink).Methods("POST")
	pathRouter.HandleFunc("/all", controller.PostAllLinks).Methods("GET")
	pathRouter.HandleFunc("/go/{path:[a-z0-9]+}", controller.GetGoLink).Methods("GET")
	return nil
}

func userRouters(router *mux.Router) error {
	userRouter, err := getSubrouter(router, "user")
	if err != nil {
		return err
	}
	userRouter.HandleFunc("/login", controller.PostLogin).Methods("POST")
	return nil
}

func getSubrouter(router *mux.Router, s string) (*mux.Router, error) {
	if s == "" {
		return nil, errors.New("prefix not founded")
	}
	return router.PathPrefix("/" + s).Subrouter(), nil
}
