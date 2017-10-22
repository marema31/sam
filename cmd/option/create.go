package option

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newOptionCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <project> <option>",
		Short: "Create a option",
		Long:  `Create the specified option in provided project.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return CreateOption(args[0], args[1])
		},
	}

	return cmd
}

func CreateOption(projectName string, optionName string) error {
	fmt.Println("create called for option:", optionName, "in project:", projectName)
	return nil
}
