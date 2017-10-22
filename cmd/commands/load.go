package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newLoadCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "load [<project> [<context>]]",
		Short: "Print the environment",
		Long: `Print all variables definitions for:
- The specified context of the provided project
- The activated context of the provided project
- The activated context of all the projects
`,
		Args: cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Load(args)
		},
	}

	return cmd
}

func Load(args []string) error {
	if len(args) == 2 {
		fmt.Println("load called for context:", args[1], "in project:", args[0])
	} else if len(args) == 1 {
		fmt.Println("load called for active context of project:", args[0])
	} else {
		fmt.Println("load called for active context of all projects")
	}
	return nil
}
