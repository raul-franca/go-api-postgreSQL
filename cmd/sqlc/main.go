package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/raul-franca/go-api-postgreSQL/internal/db"

	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:pass@localhost:5432/todos?sslmode=disable"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)
	reslt, err := queries.CreateTodo(ctx, db.CreateTodoParams{
		Name:        sql.NullString{String: "Tirar o lixo", Valid: true},
		Description: sql.NullString{String: "Recolher os lixos de casa e levar para fora.", Valid: true},
		Done:        sql.NullBool{Bool: false},
	})

	if err != nil {
		log.Fatalf("Error to create new Todo: %v", err)
	}

	fmt.Printf("Nova tarefa add, id:%d, %v \n", reslt.ID, reslt.Name.String)
	fmt.Printf("Descrição: %v \n", reslt.Description.String)
	fmt.Printf("Criada em : %v, Status: %v\n", reslt.CreatedAt.Format("02/01/06 as 15:04"), reslt.Done.Bool)

}
