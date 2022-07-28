package main

import (
	"database/sql"

	_articleHttp "tech-test/internal/article/http"
	_articleRepo "tech-test/internal/article/repository"
	_articleUsecase "tech-test/internal/article/usecase"

	_authorRepo "tech-test/internal/author/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {

	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/test?parseTime=true")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	e := echo.New()

	authorRepo := _authorRepo.NewAuthorRepository(db)
	articleRepo := _articleRepo.NewArticleRepo(db)

	articleUsecase := _articleUsecase.NewArticleUsecase(articleRepo, authorRepo)

	_articleHttp.NewArticleHandler(e, articleUsecase)

	e.Logger.Fatal(e.Start(":8000"))
}
