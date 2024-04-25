package softbase

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/spf13/cobra"
)

var Version = "0.0.1"

// SoftBase defines a SoftBase application launcher.
type SoftBase struct {
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

func (sb *SoftBase) Start(commands ...*cobra.Command) error {
    for _, cmd := range commands {
        sb.RootCmd.AddCommand(cmd)
    }

	return sb.Execute()
}

func (sb *SoftBase) Execute() error {
    // Create a new store with a degree of 2

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
    
    // core.SaveToDisk("store.db", core.Store.tree)

	// TODO, add a graceful shutdown here
	return nil
}
