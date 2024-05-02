package softbase

import (
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
    
    "github.com/LOTaher/softbase/core"
	"github.com/spf13/cobra"
)

var Version = "0.0.1"

type SoftBase struct {
	RootCmd *cobra.Command
    DB *core.Store
}

func HasDatabase() bool {
    _, err := os.Stat("softbase.gob")
    return !os.IsNotExist(err)
}

func LoadDatabase() *core.Store {
    db := core.NewStore(2)
    db.LoadFromDisk("softbase.gob")
    return db
}

func New(db *core.Store) *SoftBase {
	sb := &SoftBase{
		RootCmd: &cobra.Command{
			Use:     filepath.Base(os.Args[0]),
			Short:   "SoftBase is a key-value store backend for your next side project.",
			Version: Version,
		},
        DB: db,
	}

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

	done := make(chan bool, 1)

	go func() {
		sigch := make(chan os.Signal, 1)
		signal.Notify(sigch, os.Interrupt, syscall.SIGTERM)
		<-sigch

        if err := sb.DB.SaveToDisk("softbase.gob"); err != nil {
            panic(err)
        }

		done <- true
	}()

	go func() {
		sb.RootCmd.Execute()

		done <- true
	}()

	<-done
    
	// TODO, add a graceful shutdown here
	return nil
}
