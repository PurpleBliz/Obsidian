package main

/*
=====> Sources:
=====> migrations: https://habr.com/ru/articles/809301/  &&  https://dev.to/ynrfin/use-golang-migrate-on-docker-compose-50o5?ysclid=m9dnqzrtul804648921
=====> fiber: https://habr.com/ru/companies/otus/articles/841194/  &&  https://gofiber.io
=====> pgx: https://donchev.is/post/working-with-postgresql-in-go-using-pgx/
=====> swagger: https://docs.gofiber.io/contrib/swagger/
*/

/*
I want to create a test but, I've to go
*/

/*
swag init -d cmd
*/

//TODO: Tests need to be created
import (
	_ "Obsidian/docs"
	conf "Obsidian/internal/cfg"
	"Obsidian/internal/repository"
	"Obsidian/internal/services"
	"Obsidian/internal/transport"
)

// @title Fiber Swagger Example API
// @version 1.0
// @description This is a sample server...

// @contact.name API Support

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	cfg := conf.NewConfig()
	err := services.UpDatabase(cfg.DBConnectionString)
	if err != nil {
		panic(err)
	}

	repo := repository.NewRepository(cfg.DBConnectionString)
	taskRepo := repository.NewTaskRepository(repo)

	tr := transport.NewTransport(taskRepo)
	err = tr.App.Listen(":" + cfg.AppPort)
	if err != nil {
		panic(err)
	}
}
