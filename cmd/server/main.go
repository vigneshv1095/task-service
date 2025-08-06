package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/vigneshv1095/task-service/internal/handler"
    "github.com/vigneshv1095/task-service/internal/repository"
    "github.com/vigneshv1095/task-service/internal/service"
    _ "github.com/vigneshv1095/task-service/docs" // swagger docs
    httpSwagger "github.com/swaggo/http-swagger"
)

// @title Task Management Service API
// @version 1.0
// @description A simple Go microservice for managing tasks (CRUD), with pagination & status filtering.
// @host localhost:8080
// @BasePath /
func main() {
    repo := repository.NewInMemoryRepo()
    svc := service.NewTaskService(repo)
    h := handler.NewTaskHandler(svc)

    r := mux.NewRouter()
    h.Register(r)

    // Swagger UI endpoint
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    log.Println("Server listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
