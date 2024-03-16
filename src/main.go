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

	branches, err := services.GetListBranches(os.Getenv("INPUT_OWNER"), os.Getenv("INPUT_REPO"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Lista de branches:")
	for _, branch := range *branches {
		branchDetail, err := services.GetBranchDetails(os.Getenv("INPUT_OWNER"), os.Getenv("INPUT_REPO"), branch.Name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("----------------------------")
		message := fmt.Sprintf("Branch: %s\nAutor: %s\nData: %s\n", branchDetail.Name, branchDetail.Commit.Commit.Author.Name, branchDetail.Commit.Commit.Author.Date)
		message += fmt.Sprintf("Committer: %s\nData: %s\n", branchDetail.Commit.Commit.Committer.Name, branchDetail.Commit.Commit.Committer.Date)
		message += fmt.Sprintf("Mensagem: %s\n", branchDetail.Commit.Commit.Message)
		message += fmt.Sprintf("Total de comentários: %d\n", branchDetail.Commit.Commit.CommentCount)
		message += fmt.Sprintf("Url: %s\n", branchDetail.Commit.Commit.Url)
		fmt.Println(message)
	}
}
