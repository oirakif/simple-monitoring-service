package main

import (
	cronHandler "demo/handlers/cron"
	httpHandler "demo/handlers/http"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	cronHandler.Init(c)

	r := mux.NewRouter()
	httpHandler.Init(r)
}
