package branches

import (
	"fmt"

	"github.com/FINTLabs/fint-jsonschema-generator/common/github"
	"github.com/urfave/cli"
)

// CmdListBranches implements the `listBranches` command.
func CmdListBranches(c *cli.Context) {
	for _, b := range github.GetBranchList(c.GlobalString("owner"), c.GlobalString("repo")) {
		fmt.Println(b)
	}
}
