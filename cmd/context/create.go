package context

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newContextCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <project> <context>",
		Short: "Create a context",
		Long:  `Create the specified context in provided project.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return CreateContext(args[0], args[1])
		},
	}

	return cmd
}

func CreateContext(projectName string, contextName string) error {
	fmt.Println("create called for context:", contextName, "in project:", projectName)
	return nil
}
