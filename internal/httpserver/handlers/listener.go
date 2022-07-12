package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
)

func GoListenRutine() {
	// маршрутизация запросов обработчику
	r := chi.NewRouter()

	// зададим встроенные middleware, чтобы улучшить стабильность приложения
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// so mach information, switch ON only for debug
	if config.ConfigCLS.DebugLogger == "+" {
		r.Use(middleware.Logger)

	}
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.Compress(5, "application/json", "html/text", "text/plain", "text/html"))

	// GET requests
	r.Get("/", http.NotFound)

	// POST requests
	r.Post("/", http.NotFound)
	r.Post("/api/user/register", PostUserRegister)
	r.Post("/api/user/login", PostUserLogin)
	r.Post("/api/user/orders", PostUserOrders)

	// start listener http
	go ListenRutine(r)

}

func ListenRutine(r *chi.Mux) {
	if err := http.ListenAndServe(config.ConfigCLS.RunAddress, r); err != nil {
		config.LoggerCLS.Fatal(err.Error())
	}
}
