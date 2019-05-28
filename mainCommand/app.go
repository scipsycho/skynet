package main

import (
	"fmt"
	"log"

	root "skynet/pkg"
	"skynet/pkg/config"
	"skynet/pkg/mongo"
	"skynet/pkg/server"
)

// App forms the core struct for running the site
type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

// Initialize bootstraps the app
func (a *App) Initialize() {
	a.config = config.GetConfig()

	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatal("unable to connect to mongodb")
	}

	a.server = server.NewServer(a.config)
	a.server.CreateRoutes()
	a.server.CreateUserRouter(mongo.NewUserService(a.session, a.config.Mongo))
	a.server.CreateRecordRouter(mongo.NewRecordService(a.session, a.config.Mongo))
	a.server.CreateClaimRouter(mongo.NewClaimService(a.session, a.config.Mongo))
}

// Run starts the server
func (a *App) Run() {
	fmt.Println("Run")
	defer a.session.Close()
	a.server.Start()
}

func main() {
	a := App{}
	a.Initialize()
	a.Run()
}
