package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	config "github.com/zvovayar/yandex-praktikum-go-diploma-1/internal/config/cls"
)

var TokenAuth *jwtauth.JWTAuth

func GoListenRutine() {
	// маршрутизация запросов обработчику
	r := chi.NewRouter()

	// protected API
	r.Group(func(r chi.Router) {
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

		// JWT init

		claims := make(map[string]interface{})
		claims["user_id"] = 123
		claims["exp"] = jwtauth.ExpireIn(time.Minute * time.Duration(config.ConfigCLS.TokenTimountMinutes))
		_, tokenString, _ := TokenAuth.Encode(claims)

		config.LoggerCLS.Sugar().Debug("for debug only! sample jwt for user 123 is %s\n\n", tokenString)

		r.Use(jwtauth.Verifier(TokenAuth))
		r.Use(jwtauth.Authenticator)

		// GET requests
		r.Get("/api/user/orders", GetUserOrders)
		r.Get("/api/user/balance", GetUserBalance)
		r.Get("/api/user/balance/withdraw", GetUserBalanceWithdrawals)

		// POST requests
		r.Post("/", http.NotFound)
		r.Post("/api/user/orders", PostUserOrders)
		r.Post("/api/user/balance/withdraw", PostUserBalanceWithdraw)
	})

	// anonymous API
	r.Group(func(r chi.Router) {
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
		r.Post("/api/user/register", PostUserRegister)
		r.Post("/api/user/login", PostUserLogin)
	})

	// start listener http
	go ListenRutine(r)

}

func ListenRutine(r *chi.Mux) {
	if err := http.ListenAndServe(config.ConfigCLS.RunAddress, r); err != nil {
		config.LoggerCLS.Fatal(err.Error())
	}
}

func init() {
	TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

}
