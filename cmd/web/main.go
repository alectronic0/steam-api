package main

import (
	"fmt"
	"log"
	"steam-api/internal/config"
	"steam-api/internal/httpserver"
)

func main() {
	httpServerConfig, err := config.LoadHttpServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := httpserver.New(httpServerConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("server started on %s\n", server.Addr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
