package option

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newOptionDisableCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "disable <project> <context> <option>",
		Short: "Disable a option",
		Long:  `Disable the specified option in provided context of a project.`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return DisableOption(args[0], args[1], args[2])
		},
	}

	return cmd
}

func DisableOption(projectName string, contextName string, optionName string) error {
	fmt.Println("disable called for option:", optionName, "in context:", contextName, "of project:", projectName)
	return nil
}
