package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/srgupta5328/Employee/pkg/router"
)

const (
	PORT = ":8934"
)

func Execute() {
	fmt.Println("This is an Employee Directory Application")
	r := router.NewRouter()
	log.Fatal(http.ListenAndServe(PORT, r))

}
