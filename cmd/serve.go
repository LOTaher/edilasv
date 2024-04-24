package cmd

import (
	"github.com/LOTaher/softbase/apis"
	"github.com/spf13/cobra"
    "github.com/LOTaher/softbase/core"
)

func Serve(db *core.Store) *cobra.Command {
    var allowedOrigins []string
    var httpAddr string

    command := &cobra.Command{
        Use:  "serve",
        Short: "Start the server (defaults to port 127.0.0.1:1404 if no domain is specified)",
        RunE: func(cmd *cobra.Command, args []string) error {

            if httpAddr == "" {
                httpAddr = "127.0.0.1:1404"
            }
            
            _, err := apis.Serve(apis.ServeConfig{
                AllowedOrigins: allowedOrigins,
                HttpAddr:       httpAddr,
                DB:             db,
            })
            if err != nil {
                return err
            }

            return nil
        },
    }

    command.Flags().StringSliceVar(
        &allowedOrigins,
        "origins",
        []string{"*"},
        "Allowed origins for CORS requests",
    )

    command.Flags().StringVar(
        &httpAddr,
        "http",
        "",
        "HTTP service address",
    )

    return command

}

    

    
