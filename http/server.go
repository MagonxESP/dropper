package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ginserver "github.com/go-oauth2/gin-server"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	"log"
	"os"
)

func OAuth2Manager() *manage.Manager {
	manager := manage.NewDefaultManager()

	if _, err := os.ReadDir("storage"); os.IsNotExist(err) {
		if err := os.Mkdir("storage", 0755); err != nil {
			panic(err)
		}
	}

	manager.MustTokenStorage(store.NewFileTokenStore("storage/oauth2.db"))
	clientStore := store.NewClientStore()
	err := clientStore.Set(os.Getenv("DROPPER_OAUTH2_DEFAULT_CLIENT_ID"), &models.Client{
		ID:     os.Getenv("DROPPER_OAUTH2_DEFAULT_CLIENT_ID"),
		Secret: os.Getenv("DROPPER_OAUTH2_DEFAULT_CLIENT_SECRET"),
		Domain: os.Getenv("DROPPER_OAUTH2_DEFAULT_CLIENT_DOMAIN"),
	})

	if err != nil {
		panic(err)
	}

	manager.MapClientStorage(clientStore)
	return manager
}

func GetHttpServerPort() string {
	port := os.Getenv("DROPPER_HTTP_PORT")

	if port == "" {
		port = "8080"
	}

	return port
}

func StartHttpServer() {
	manager := OAuth2Manager()
	ginserver.InitServer(manager)
	ginserver.SetAllowGetAccessRequest(true)
	ginserver.SetClientInfoHandler(server.ClientFormHandler)

	engine := gin.Default()

	RegisterRoutes(engine)
	err := engine.Run(fmt.Sprintf(":%s", GetHttpServerPort()))

	if err != nil {
		log.Fatal(err)
	}
}
