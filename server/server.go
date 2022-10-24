package server

import (
	"fmt"
	"os"

	"finall/config"

	"github.com/gin-gonic/gin"
)

func Start() error {
	db, err := config.New()
	if err != nil {
		return err
	}

	r := gin.Default()
	NewRouter(r, db)

	r.Use(gin.Recovery())

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	r.Run(fmt.Sprintf(":%s", port))
	return nil
}
