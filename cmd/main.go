package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/devfullcycle/pfa-go/internal/order/infra/database"
	"github.com/devfullcycle/pfa-go/internal/order/usecase"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	var err error

	dsn := "root:root@tcp(mysql:3306)/orders"

	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Printf("MySQL not ready, retrying in 5 seconds (%d/10)\n", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize the database
	err = database.InitializeDB(db)
	if err != nil {
		panic(err)
	}

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	input := usecase.OrderInputDTO{
		ID:    "1234",
		Price: 100,
		Tax:   10,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output.FinalPrice)
}
