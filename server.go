package api

import (
	"api/auto"
	"api/config"
	"api/router"
	"fmt"
	"log"
	"net/http"
)

// Run is the main function for setting up the server
func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("Listening on [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen (port int) {
	r := router.NEW()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}