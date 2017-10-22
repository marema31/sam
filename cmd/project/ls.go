package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newProjectListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List all projects",
		Long:  `List all projects.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListProject()
		},
	}

	return cmd
}

func ListProject() error {
	fmt.Println("ls called for projects")
	return nil
}
