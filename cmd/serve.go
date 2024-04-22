package cmd

import (
	"github.com/LOTaher/softbase/apis"
	"github.com/LOTaher/softbase/core"
	"github.com/spf13/cobra"
)

func Serve(app core.App) *cobra.Command {
    var allowedOrigins []string
    var httpAddr string

    command := &cobra.Command{
        Use:  "serve",
        Short: "Start the server (defaults to port 127.0.0.1:0414 if no domain is specified)",
        RunE: func(cmd *cobra.Command, args []string) error {

            if httpAddr == "" {
                httpAddr = "127.0.0.1:0414"
            }
            
            _, err := apis.Serve(apis.ServeConfig{
                AllowedOrigins: allowedOrigins,
                HttpAddr:       httpAddr,
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

    

    
