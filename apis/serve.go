package apis

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
    "github.com/LOTaher/softbase/core"
	// "github.com/go-chi/cors"
)

type ServeConfig struct {
	AllowedOrigins []string
	HttpAddr       string
	DB             *core.Store
}

func Serve(config ServeConfig) (*http.Server, error) {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = []string{"*"}
	}

	schema := "http"

	router, err := InitAPI(config.DB)
	if err != nil {
		return nil, err
	}
	//
	// router.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: config.AllowedOrigins,
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300,
	// }))

	server := &http.Server{
		Addr:    config.HttpAddr,
		Handler: router,
	}

	date := new(strings.Builder)
	bold := color.New(color.Bold).Add(color.FgCyan)

	fmt.Print("∘˙○˚.•\n")
	bold.Printf(
		"%s Server started at %s\n",
		strings.TrimSpace(date.String()),
		color.CyanString("%s://%s", schema, server.Addr),
	)
	regular := color.New()
	regular.Printf("├─ REST API: %s\n", color.CyanString("%s://%s/api/", schema, server.Addr))
	regular.Printf("└─ Admin Panel: %s\n", color.CyanString("%s://%s/", schema, server.Addr))
	fmt.Println("∘˙○˚.•")

	return server, server.ListenAndServe()
}
