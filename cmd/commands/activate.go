package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newActivateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "activate <project> <context>",
		Short: "Activate a context",
		Long: `Activate the specified context in provided project.
The context are mutually exclusive within a project, the 
previous activated context will be deactivated`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return Activate(args[0], args[1])
		},
	}

	return cmd
}

func Activate(projectName string, contextName string) error {
	fmt.Println("activate called for context:", contextName, "in project:", projectName)
	return nil
}
