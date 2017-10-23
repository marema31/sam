package commands

import (
	"fmt"

	"github.com/marema31/sam/manage"
	"github.com/spf13/cobra"
)

func newInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Init SAM path",
		Long: `Create the folder used by SAM 
project contexts management.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Init()
		},
	}

	return cmd
}

func Init() error {
	manage.CreateSamPath()
	fmt.Println("Initialization completed !")
	return nil
}
