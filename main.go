package main

import (
	"contacts-0/httpRouter"
	"fmt"
	"net/http"
)

func main() {
	router := httpRouter.CreateNewRouter()

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Working as expected!")
	})

	router.Run()
}
