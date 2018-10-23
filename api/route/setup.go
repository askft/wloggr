package route

import (
	"net/http"
	"time"

	"wloggr/api/util"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/unrolled/secure"
)

// SetupRouter creates and sets up a chi router with some essential
// middleware (logging, security, CORS, and more).
func SetupRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	setupSecure(r)
	setupCORS(r)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello! This is a public API for use with wloggr.com.\n"))
	})

	r.Route("/api", func(r chi.Router) {
		r.Mount("/workout", WorkoutRoutes())
		r.Mount("/user", UserRoutes())
	})

	return r
}

func setupSecure(r *chi.Mux) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP
	secureMiddleware := secure.New(secure.Options{
		// IsDevelopment disables AllowedHosts, SSLRedirect, STS header, HPKP header
		IsDevelopment:         util.Config.IsDevelopment,
		AllowedHosts:          []string{util.Config.DomainName, "ssl." + util.Config.DomainName},
		HostsProxyHeaders:     []string{"X-Forwarded-Host"},
		SSLRedirect:           true,
		SSLHost:               "ssl." + util.Config.DomainName,
		SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		STSPreload:            true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "script-src $NONCE",
	})
	r.Use(secureMiddleware.Handler)
}

func setupCORS(r *chi.Mux) {
	// https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
}
