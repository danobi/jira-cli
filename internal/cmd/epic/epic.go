package epic

import (
	"github.com/spf13/cobra"

	"github.com/ankitpokhrel/jira-cli/internal/cmd/epic/create"
	"github.com/ankitpokhrel/jira-cli/internal/cmd/epic/list"
)

const helpText = `Epic manage epics in a given project. See available commands below.`

// NewCmdEpic is an epic command.
func NewCmdEpic() *cobra.Command {
	cmd := cobra.Command{
		Use:     "epic [ISSUE KEY]",
		Short:   "Epic manage epics in a project.",
		Long:    helpText,
		Aliases: []string{"epics"},
		RunE:    epic,
	}

	lc := list.NewCmdList()
	cc := create.NewCmdCreate()

	cmd.AddCommand(lc, cc)

	list.SetFlags(lc)
	create.SetFlags(cc)

	return &cmd
}

func epic(cmd *cobra.Command, _ []string) error {
	return cmd.Help()
}
