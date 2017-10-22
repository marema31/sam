package context

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newContextRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm <project> <context>",
		Short: "Remove a context",
		Long:  `Remove the specified context.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RemoveContext(args[0], args[1])
		},
	}

	return cmd
}

func RemoveContext(projectName string, contextName string) error {
	fmt.Println("rm called for context:", contextName, "in project:", projectName)
	return nil
}
