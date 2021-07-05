package main

import (
	"database/sql"
	"fmt"
	"github.com/joferreira/codebank/domain"
	"github.com/joferreira/codebank/infrastructure/repository"
	"github.com/joferreira/codebank/usecase"
	"log"
	_ "github.com/lib/pq"
)

func main(){
	db := setupDB()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Josemar"
	cc.ExpirationMonth = 7
	cc.ExpirationYear = 2021
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}

}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDB() *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}

	return db
}

