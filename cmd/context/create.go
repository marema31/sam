package context

import (
	"fmt"

	"github.com/marema31/sam/manage"
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

func CreateContext(project string, context string) error {
	err := manage.CreateContext(project, context)
	if err == nil {
		fmt.Println("Context", context, "of project", project, "created !")
	}
	return err
}
