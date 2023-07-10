package main

import (
	"exercise/router"
	"exercise/utils"
	"net/http"
	"time"
)

func main() {
	router := router.InitRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	utils.InitMysql()
	s.ListenAndServe()
}
