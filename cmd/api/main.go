package main

import (
	"database/sql"
	"fmt"
	"go_searcher/http/echo"
	"go_searcher/internal/api"
	"go_searcher/participant"
	"go_searcher/participant/sqlite"
	programapocalipse "go_searcher/program_apocalipse"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./inscriptions.db")
	if err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	repo := sqlite.NewSqlite(db)
	pService := participant.NewService(repo)
	cService := programapocalipse.NewService()
	h := echo.Handlers(pService, cService)
	err = api.Start("8080", h)
	if err != nil {
		fmt.Printf("error running api %s", err)
	}

}
