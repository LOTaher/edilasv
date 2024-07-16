package cmd

import (
	"github.com/LOTaher/softbase/apis"
	"github.com/LOTaher/softbase/core"
	"github.com/spf13/cobra"
)

func Serve(db *core.Store) *cobra.Command {
	var allowedOrigins []string
	var httpAddr string
	var httpsAddr string

	command := &cobra.Command{
		Use:   "serve",
		Short: "Start the server (defaults to port 127.0.0.1:1404 if no domain is specified)",
		RunE: func(cmd *cobra.Command, args []string) error {

			// set the default listener addresses if at least one domain is not specified
			if len(args) > 0 {
				if httpAddr == "" {
					httpAddr = "0.0.0.0:80"
				}

				if httpsAddr == "" {
					httpsAddr = "0.0.0.0:443"
				}
			} else {
				if httpAddr == "" {
					httpAddr = "127.0.0.1:1404"
				}
			}

			_, err := apis.Serve(apis.ServeConfig{
				AllowedOrigins: allowedOrigins,
				HttpAddr:       httpAddr,
                HttpsAddr:      httpsAddr,
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

	command.Flags().StringVar(
		&httpsAddr,
		"https",
		"",
		"HTTPS service address",
	)

	return command
}
