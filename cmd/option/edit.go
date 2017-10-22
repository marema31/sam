package option

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newOptionEditCommand() *cobra.Command {
	var editService bool

	cmd := &cobra.Command{
		Use:   "edit <project> <option>",
		Short: "Edit a option",
		Long: `Edit the specified option.
Without flags, environement will be edited.`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return EditOption(args[0], args[1], editService)
		},
	}
	cmd.Flags().BoolVarP(&editService, "services", "s", false, "Edit the services")

	return cmd
}

func EditOption(projectName string, optionName string, editService bool) error {
	fmt.Println("edit called for option:", optionName, "in project:", projectName, "with service:", editService)

	return nil
}
