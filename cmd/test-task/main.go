package main

import (
	"log"

	"github.com/wehw93/test_task_Junior_go_developer/internal/pkg/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal(err)
	}
	err = a.Run()
	if err != nil {
		log.Fatal(err)
	}

}
