package main

import (
	"database/sql"
	"fmt"
	"os"

	_articleHttp "tech-test/internal/article/http"
	_articleRepo "tech-test/internal/article/repository"
	_articleUsecase "tech-test/internal/article/usecase"

	_authorRepo "tech-test/internal/author/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	username := os.Getenv("DB_USERNAME")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "toor"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "127.0.0.1"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "test"
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, name))

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	e := echo.New()

	authorRepo := _authorRepo.NewAuthorRepository(db)
	articleRepo := _articleRepo.NewArticleRepo(db)

	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, authorRepo)

	_articleHttp.NewArticleHandler(e, articleUsecase)

	e.Logger.Fatal(e.Start(":8000"))
}
