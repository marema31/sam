package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newPsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ps [<project> [<context>]]",
		Short: "",
		Long: `Print project, activated context and enabled options of:
- The specified context of the provided project
- The activated context of the provided project
- The activated context of all the projects
`,
		Args: cobra.MaximumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Ps(args)
		},
	}

	return cmd
}

func Ps(args []string) error {
	if len(args) == 2 {
		fmt.Println("ps called for context:", args[1], "in project:", args[0])
	} else if len(args) == 1 {
		fmt.Println("ps called for active context of project:", args[0])
	} else {
		fmt.Println("ps called for active context of all projects")
	}
	return nil
}
