package cmds

import (
	"fmt"
	"strings"

	"github.com/appscode/go/log"
	"github.com/appscode/guard/appscode"
	"github.com/appscode/guard/github"
	"github.com/appscode/guard/gitlab"
	"github.com/appscode/guard/google"
	"github.com/appscode/guard/ldap"
	"github.com/appscode/guard/server"
	"github.com/spf13/cobra"
)

type tokenOptions struct {
	Org  string
	LDAP ldap.TokenOptions
}

func NewCmdGetToken() *cobra.Command {
	opts := tokenOptions{}

	cmd := &cobra.Command{
		Use:               "token",
		Short:             fmt.Sprintf("Get tokens for %v", server.SupportedOrgPrintForm()),
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			opts.Org = strings.ToLower(opts.Org)
			switch opts.Org {
			case github.OrgType:
				github.IssueToken()
				return
			case gitlab.OrgType:
				gitlab.IssueToken()
				return
			case google.OrgType:
				err := google.IssueToken()
				if err != nil {
					log.Fatal(err)
				}
				return
			case appscode.OrgType:
				err := appscode.IssueToken()
				if err != nil {
					log.Fatal(err)
				}
				return
			case ldap.OrgType:
				err := opts.LDAP.IssueToken()
				if err != nil {
					log.Fatal("For LDAP:", err)
				}
				return
			case "":
				log.Fatalln("Missing organization name. Set flag -o Google|Github.")
			default:
				log.Fatalf("Unknown organization %s.", opts.Org)
			}
		},
	}

	cmd.Flags().StringVarP(&opts.Org, "organization", "o", opts.Org, fmt.Sprintf("Name of Organization (%v).", server.SupportedOrgPrintForm()))
	opts.LDAP.AddFlags(cmd.Flags())
	return cmd
}
