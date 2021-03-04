package main

import (
	"gb-go-architecture/lesson-2/shop_new/notification"
	"gb-go-architecture/lesson-2/shop_new/repository"
	"gb-go-architecture/lesson-2/shop_new/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	TgChatID   int64
	TgTokenStr string
	SMTPHost   string
	SMTPPort   string
	SMTPUser   string
	SMTPPass   string
}

func NewRouter(s *server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/items", s.listItemHandler).Methods("GET")
	router.HandleFunc("/items", s.createItemHandler).Methods("POST")
	router.HandleFunc("/items/{id}", s.getItemHandler).Methods("GET")
	router.HandleFunc("/items/{id}", s.deleteItemHandler).Methods("DELETE")
	router.HandleFunc("/items/{id}", s.updateItemHandler).Methods("PUT")

	router.HandleFunc("/orders", s.listOrdersHandler).Methods("GET")
	router.HandleFunc("/orders", s.createOrderHandler).Methods("POST")

	return router
}

func main() {
	var appConf Config
	if err := envconfig.Process("myshop", &appConf); err != nil {
		log.Fatalf("Cannot load config, error: %s", err.Error())
	}

	staffNotif, err := notification.NewTelegramBot(appConf.TgTokenStr, appConf.TgChatID)
	if err != nil {
		log.Fatal(err)
	}
	userNotif, err := notification.NewSMTPClient(
		appConf.SMTPUser, appConf.SMTPPass, appConf.SMTPHost, appConf.SMTPPort,
	)
	if err != nil {
		log.Fatal(err)
	}

	rep := repository.NewMapDB()
	service := service.NewService(rep, staffNotif, userNotif)
	s := &server{
		service: service,
		rep:     rep,
	}

	srv := &http.Server{
		Addr:    ":8081",
		Handler: NewRouter(s),
	}
	log.Fatal(srv.ListenAndServe())
}
