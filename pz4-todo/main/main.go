package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"

	"example.com/pz4-todo/internal/task"
	myMW "example.com/pz4-todo/pkg/middleware"
)

func main() {
	repo := task.NewRepo()
	inFileRepo := task.NewFileRepo("tasks.json")
	h := task.NewHandler(repo)
	handler := task.NewHandler(inFileRepo)

	// Создаем роутер
	r := chi.NewRouter()

	// Просто подкидываем MVs
	r.Use(chimw.RequestID)
	r.Use(chimw.Recoverer)
	r.Use(myMW.Logger)
	r.Use(myMW.SimpleCORS)

	// Ручка для проверки состояния на уровне хоста
	r.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		task.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	r.Route("/api/v1", func(api chi.Router) {
		api.Mount("/tasks", h.Routes())
		api.Get("/health", h.Health)
	})

	r.Route("/api/v2", func(api chi.Router) {
		api.Mount("/tasks", handler.Routes()) // Для v2 используется inFileRepo
		api.Get("/health", h.Health)
	})

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
