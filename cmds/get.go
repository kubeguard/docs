package cmds

import (
	"github.com/spf13/cobra"
)

func NewCmdGet() *cobra.Command {
	cmd := &cobra.Command{
		Use:               "get",
		Short:             `Get PKI`,
		DisableAutoGenTag: true,
	}
	cmd.AddCommand(NewCmdGetWebhookConfig())
	cmd.AddCommand(NewCmdGetToken())
	cmd.AddCommand(NewCmdInstaller())
	return cmd
}
