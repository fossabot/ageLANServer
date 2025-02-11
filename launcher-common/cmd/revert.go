package cmd

import (
	"fmt"
	commonCmd "github.com/luskaner/ageLANServer/common/cmd"
	launcherCommon "github.com/luskaner/ageLANServer/launcher-common"
	"github.com/spf13/cobra"
)

var RemoveLocalCert bool
var UnmapIPs bool
var UnmapCDN bool
var RemoveAll bool

func InitRevert(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(
		&UnmapIPs,
		"ip",
		"i",
		false,
		fmt.Sprintf("Remove the IP mappings from the local DNS server"),
	)
	cmd.Flags().BoolVarP(
		&UnmapCDN,
		"CDN",
		"c",
		false,
		fmt.Sprintf("Remove the mappings from the local DNS server that resolve %s", launcherCommon.CDNIP),
	)
	cmd.Flags().BoolVarP(
		&RemoveLocalCert,
		"localCert",
		"l",
		false,
		"Remove the certificate from the local machine's trusted root store",
	)
	cmd.Flags().BoolVarP(
		&RemoveAll,
		"all",
		"a",
		false,
		"Removes all configuration. Equivalent to the rest of the flags being set without fail-fast.",
	)
	commonCmd.GameVarCommand(cmd.Flags(), &GameId)
}
