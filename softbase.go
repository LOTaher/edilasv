package softbase

import (
	"os"
    "os/signal"
	"path/filepath"
    "syscall"

	"github.com/LOTaher/softbase/cmd"
	"github.com/LOTaher/softbase/core"
	"github.com/spf13/cobra"
)

var Version = "0.0.1"

type appWrapper struct {
	core.App
}

// SoftBase defines a SoftBase application launcher.
type SoftBase struct {
	*appWrapper
	RootCmd *cobra.Command
}

// Creates a new SoftBase instance with the default configuration.
func New() *SoftBase {
	sb := &SoftBase{
		RootCmd: &cobra.Command{
			Use:     filepath.Base(os.Args[0]),
			Short:   "SoftBase is a key-value store backend for your next side project.",
			Version: Version,
		},
	}

	// hide the default help command, only allowing `--help` flag.
	sb.RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})

	return sb
}

func (sb *SoftBase) Start() error {
	sb.RootCmd.AddCommand(cmd.Serve(sb))

	return sb.Execute()
}

func (sb *SoftBase) Execute() error {
	done := make(chan bool, 1)

	// listen for interrupt signal to gracefully shutdown the application
	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch

		done <- true
	}()

	// execute the root command
	go func() {
		sb.RootCmd.Execute()

		done <- true
	}()

	<-done
    
    // TODO, add a graceful shutdown here
    return nil
}
