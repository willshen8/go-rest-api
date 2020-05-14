package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/willshen8/go-rest-api/pkg/puppy"
)

func main() {
	router := puppy.SetupRouter()
	puppy.SetupLostPuppyRoutes(router)
	logrus.Fatal(http.ListenAndServe(":8888", router))
}
