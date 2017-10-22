package option

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newOptionEnableCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable <project> <context> <option>",
		Short: "Enable a option",
		Long:  `Enable the specified option in provided context of a project.`,
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			return EnableOption(args[0], args[1], args[2])
		},
	}

	return cmd
}

func EnableOption(projectName string, contextName string, optionName string) error {
	fmt.Println("enable called for option:", optionName, "in context:", contextName, "of project:", projectName)
	return nil
}
