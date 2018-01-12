package utils



import (
	"gopkg.in/urfave/cli.v1"
	"path/filepath"
	"os"
	"strings"
	"github.com/ofbank_wallet/OFBANK_WALLET/node"
	"github.com/ofbank_wallet/OFBANK_WALLET/common"
	"math/big"
	"github.com/ofbank_wallet/OFBANK_WALLET/params"
	"fmt"
)

func NewApp(gitCommit, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = ""
	//app.Authors = nil
	app.Email = ""
	//ARVIN-VERSION
	app.Version = "version1"
	if gitCommit != "" {
		app.Version += "-" + gitCommit[:8]
	}
	app.Usage = usage
	return app
}


var(

	DataDirFlag = DirectoryFlag{
		Name:  "datadir",
		Usage: "Data directory for the databases and keystore",
		Value: DirectoryString{node.DefaultDataDir()},
	}

// ATM the url is left to the user and deployment to
JSpathFlag = cli.StringFlag{
Name:  "jspath",
Usage: "JavaScript root path for `loadScript`",
Value: ".",
}
	ExecFlag = cli.StringFlag{
		Name:  "exec",
		Usage: "Execute JavaScript statement",
	}


	PreloadJSFlag = cli.StringFlag{
		Name:  "preload",
		Usage: "Comma separated list of JavaScript files to preload into the console",
	}
	TargetGasLimitFlag = cli.Uint64Flag{
		Name:  "targetgaslimit",
		Usage: "Target gas limit sets the artificial target gas floor for the blocks to mine",
		Value: params.GenesisGasLimit.Uint64(),
	}

)


func MigrateFlags(action func(ctx *cli.Context) error) func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		for _, name := range ctx.FlagNames() {
			if ctx.IsSet(name) {
				ctx.GlobalSet(name, ctx.String(name))
			}
		}
		return action(ctx)
	}
}


func MakeDataDir(ctx *cli.Context) string {

	if path := ctx.GlobalString(DataDirFlag.Name); path != "" {
       fmt.Println("datadir: ",path)
		return path
	}

	Fatalf("Cannot determine default data directory, please set manually (--datadir)")
	return ""
}

func MakeConsolePreloads(ctx *cli.Context) []string {
	// Skip preloading if there's nothing to preload
	if ctx.GlobalString(PreloadJSFlag.Name) == "" {
		return nil
	}
	// Otherwise resolve absolute paths and return them
	preloads := []string{}

	assets := ctx.GlobalString(JSpathFlag.Name)
	for _, file := range strings.Split(ctx.GlobalString(PreloadJSFlag.Name), ",") {
		preloads = append(preloads, common.AbsolutePath(assets, strings.TrimSpace(file)))
	}
	return preloads
}
// SetupNetwork configures the system for either the main net or some test network.
func SetupNetwork(ctx *cli.Context) {
	// TODO(fjl): move target gas limit into config
	params.TargetGasLimit = new(big.Int).SetUint64(ctx.GlobalUint64(TargetGasLimitFlag.Name))
}

