package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matheusrosmaninho/github-branches-list/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	branches, err := services.GetListBranches(os.Getenv("ACCOUNT_OWNER"), os.Getenv("REPO_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lista de branches:")
	for _, branch := range *branches {
		fmt.Println(branch)
	}
}
