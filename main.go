package main

import (
	"database/sql"
	"fmt"

	"github.com/DaviAugustoMS/go-internsivo/internal/infra/database"
	"github.com/DaviAugustoMS/go-internsivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}

	output := usecase.OrderOutput{
		ID:    input.ID,
		Price: input.Price,
		Tax:   input.Tax,
	}

	result, err := uc.Execute(output)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
