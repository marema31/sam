package option

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newOptionListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls <project>",
		Short: "List all options of project",
		Long:  `List all options of provided project.`,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return ListOption(args[0])
		},
	}

	return cmd
}

func ListOption(projectName string) error {
	fmt.Println("ls called for option for project:", projectName)
	return nil
}
