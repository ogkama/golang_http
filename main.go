package main

import (
	"flag"
	"http_server/repository/db_storage"
	"http_server/repository/ram_storage"
	"http_server/usecases/service"
	"http_server/usecases/task_service"
	"log"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"

	"http_server/api/http"
	_ "http_server/docs"
	pkgHttp "http_server/pkg/http"
)

// @title My API
// @version 1.0
// @description This is a sample server.

// @host localhost:8080
// @BasePath /
func main() {
	addr := flag.String("addr", ":8080", "address for http server")

	objectRepo := ram_storage.NewObject()
	objectService := service.NewObject(objectRepo)
	objectHandlers := http.NewHandler(objectService)

	taskRepo := db_storage.NewObject()
	taskService := task_service.NewObject(taskRepo)
	taskHandlers := http.NewTaskHandler(taskService)

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.WrapHandler)

	objectHandlers.WithObjectHandlers(r)
	taskHandlers.WithTaskHandlers(r)

	log.Printf("Starting server on %s", *addr)
	if err := pkgHttp.CreateAndRunServer(r, *addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
