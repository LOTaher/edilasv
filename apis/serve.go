package apis

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
    "github.com/LOTaher/softbase/core"
)

type ServeConfig struct {
    // AllowedOrigins is an optional list of CORS origins (default: "*")
	AllowedOrigins []string

    // HttpAddr is the address to listen on for HTTP requests (eg. `127.0.0.1:1404`)
	HttpAddr       string

    // HttpsAddr is the address to listen on for HTTPS requests (eg. `127.0.0.1:443`)
    HttpsAddr      string

    // DB is the database store
	DB             *core.Store
}

func Serve(config ServeConfig) (*http.Server, error) {
	if len(config.AllowedOrigins) == 0 {
		config.AllowedOrigins = []string{"*"}
	}

	schema := "http"
    if config.HttpsAddr != "" {
        schema = "https"
    }

    mainAddr := config.HttpAddr
    if config.HttpsAddr != "" {
        mainAddr = config.HttpsAddr
    }

    // initialize the API
	router, err := InitAPI(config.DB, config)
	if err != nil {
		return nil, err
	}

	server := &http.Server{
		Addr:    mainAddr,
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
