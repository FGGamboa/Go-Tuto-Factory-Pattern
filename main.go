package main

import (
	"fmt"
	"tuto-factory-pattern/configuration"
	"tuto-factory-pattern/repository"
)

func main() {
	config := &configuration.Configuration{
		Engine: "sqlserver",
		Host:   "localhost",
	}

	repo, err := repository.New(config)
	if err != nil {
		panic(err)
	}

	err = repo.Save("")
	if err != nil {
		panic(err)
	}

	data := repo.Find(1)
	fmt.Println(data)
}
