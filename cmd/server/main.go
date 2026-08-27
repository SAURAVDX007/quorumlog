package main

import (
	"log"

	"github.com/SAURAVDX007/quorumlog/internal/server"
)

func main(){
	srv := server.NewHttpServer(":8080")
	log.Fatal(srv.ListenAndServe())
}