package cmd

import (
	"fmt"
	"github.com/luskaner/ageLANServer/common"
	"github.com/luskaner/ageLANServer/common/cmd"
	"github.com/luskaner/ageLANServer/server-genCert/internal"
	"github.com/spf13/cobra"
	"os"
	"path"
	"path/filepath"
)

var replace bool
var gameId string
var Version string

var (
	rootCmd = &cobra.Command{
		Use:   filepath.Base(os.Args[0]),
		Short: "genCert generates a self-signed certificate",
		Run: func(_ *cobra.Command, _ []string) {
			exe, _ := os.Executable()
			serverExe := path.Join(filepath.Dir(filepath.Dir(exe)), common.GetExeFileName(true, common.Server))
			serverFolder := common.CertificatePairFolder(serverExe)
			if serverFolder == "" {
				fmt.Println("Failed to determine certificate pair folder")
				os.Exit(internal.ErrCertDirectory)
			}
			if !replace {
				if exists, _, _ := common.CertificatePair(gameId, serverExe); exists {
					fmt.Println("Already have certificate pair and force is false, set force to true or delete it manually.")
					os.Exit(internal.ErrCertCreateExisting)
				}
			}
			if !internal.GenerateCertificatePair(gameId, serverFolder) {
				fmt.Println("Could not generate certificate pair.")
				os.Exit(internal.ErrCertCreate)
			} else {
				fmt.Println("Certificate pair generated successfully.")
			}
		},
	}
)

func Execute() error {
	rootCmd.Version = Version
	rootCmd.PersistentFlags().BoolVarP(&replace, "replace", "r", false, "Overwrite existing certificate pair.")
	cmd.GameVarCommand(rootCmd.PersistentFlags(), &gameId)
	return rootCmd.Execute()
}
