package cmds

import (
	"crypto/x509"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/appscode/go/log"
	"github.com/appscode/go/term"
	"github.com/appscode/kutil/tools/certstore"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"k8s.io/client-go/util/cert"
)

func NewCmdInitClient() *cobra.Command {
	var org string
	cmd := &cobra.Command{
		Use:               "client",
		Short:             "Generate client certificate pair",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalln("Missing client name.")
			}
			if len(args) > 1 {
				log.Fatalln("Multiple client name found.")
			}

			cfg := cert.Config{
				CommonName: args[0],
				Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
			}

			org = strings.ToLower(org)
			switch org {
			case "github":
				cfg.Organization = []string{"Github"}
			case "google":
				cfg.Organization = []string{"Google"}
			case "appscode":
				cfg.Organization = []string{"Appscode"}
			case "":
				log.Fatalln("Missing organization name. Set flag -o Google|Github.")
			default:
				log.Fatalf("Unknown organization %s.", org)
			}

			store, err := certstore.NewCertStore(afero.NewOsFs(), filepath.Join(rootDir, "pki"), cfg.Organization...)
			if err != nil {
				log.Fatalf("Failed to create certificate store. Reason: %v.", err)
			}
			if store.IsExists(filename(cfg)) {
				if !term.Ask(fmt.Sprintf("Client certificate found at %s. Do you want to overwrite?", store.Location()), false) {
					os.Exit(1)
				}
			}

			if err = store.LoadCA(); err != nil {
				log.Fatalf("Failed to load ca certificate. Reason: %v.", err)
			}

			crt, key, err := store.NewServerCertPair(cfg.CommonName, cfg.AltNames)
			if err != nil {
				log.Fatalf("Failed to generate certificate pair. Reason: %v.", err)
			}
			err = store.WriteBytes(filename(cfg), crt, key)
			if err != nil {
				log.Fatalf("Failed to init client certificate pair. Reason: %v.", err)
			}
			term.Successln("Wrote client certificates in ", store.Location())
		},
	}

	cmd.Flags().StringVar(&rootDir, "pki-dir", rootDir, "Path to directory where pki files are stored.")
	cmd.Flags().StringVarP(&org, "organization", "o", org, "Name of Organization (Github or Google).")
	return cmd
}
