package main

import (
	"log"
	"os"

	"github.com/luzhkovn/go1fl-sprint6-final-tpl/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "INFO: ", log.LstdFlags)
	srv := server.NewServer(logger)
	err := srv.Start()
	if err != nil {
		logger.Fatal(err)
	}
}
