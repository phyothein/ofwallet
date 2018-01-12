package main

import (
	"github.com/ofbank_wallet/OFBANK_WALLET/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"os"
	"time"
	"runtime"
	"github.com/ofbank_wallet/OFBANK_WALLET/internal/debug"
	"github.com/ofbank_wallet/OFBANK_WALLET/metrics"
	"github.com/ofbank_wallet/OFBANK_WALLET/console"
	 "github.com/ofbank_wallet/OFBANK_WALLET/node"
	"sort"
)


var(
	gitCommit = ""
	app = utils.NewApp(gitCommit, "the go-ethereum command line interface")

	nodeFlags = []cli.Flag{
		utils.DataDirFlag,
	}


)


func init() {

	app.Action = geth

	app.Commands = []cli.Command{
		consoleCommand,
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Flags = append(app.Flags, consoleFlags...)
	app.Flags = append(app.Flags, nodeFlags...)
	app.Before = func(ctx *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		if err := debug.Setup(ctx); err != nil {
			return err
		}
		// Start system runtime metrics collection
		go metrics.CollectProcessMetrics(3 * time.Second)

		utils.SetupNetwork(ctx)
		return nil
	}
	app.After = func(ctx *cli.Context) error {
		debug.Exit()
		console.Stdin.Close() // Resets terminal mode.
		return nil
	}
}


func geth(ctx *cli.Context) error {
	start,err:=node.NewNode()
	if err!=nil{
		return err
	}

	start.Wait()
	return nil
}


func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}