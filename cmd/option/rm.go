package option

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newOptionRemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm <project> <option>",
		Short: "Remove a option",
		Long:  `Remove the specified option.`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return RemoveOption(args[0], args[1])
		},
	}

	return cmd
}

func RemoveOption(projectName string, optionName string) error {
	fmt.Println("rm called for option:", optionName, "in project:", projectName)
	return nil
}
