package commands

import (
	"fmt"
	"net/rpc"

	"github.com/spf13/cobra"
)

// memberDemoteCmd is used to demote a designated member node in the cluster
var memberDemoteCmd = &cobra.Command{
	Use:   "demote",
	Short: "Advises a node to 'demote'",
	Long:  ``,

	Run: memberDemote,
}

// memberDemote demotes the designated member node
func memberDemote(ccmd *cobra.Command, args []string) {

	// create an RPC client that will connect to the designated node
	client, err := rpc.Dial("tcp", fmt.Sprintf("%s:%s", fHost, fPort))
	if err != nil {
		fmt.Printf("[commands/memberDemote] rpc.Dial() failed - %s\n", err.Error())
		return
	}
	defer client.Close()

	fmt.Printf("advising '%s' to demote...\n", fHost)

	// issue a demote to the designated node
	if err := client.Call("Status.Demote", "", nil); err != nil {
		fmt.Printf("[commands/memberDemote] client.Call() failed - %s\n", err.Error())
	}
}
