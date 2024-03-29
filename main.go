package main

import (
	"time"

	route "github.com/jordanlanch/ionix-test/api/route"
	"github.com/jordanlanch/ionix-test/bootstrap"
)

func main() {

	app := bootstrap.App(".env")

	env := app.Env

	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := route.Setup(env, timeout, app.Postgresql)

	router.Run(env.ServerAddress)
}
