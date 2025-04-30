package main

import (
	"fooservice/internal/data"
	"fooservice/internal/server"
	"github.com/timpellison/flaggy/client"
)

func main() {
	run()
}

func run() {
	cl := client.NewFlaggyClient("http://localhost:8088")
	ds := data.NewAccountService("new-retriever", *cl)
	s := server.NewService(ds)
	s.Start()
}
