package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/ofbank_wallet/OFBANK_WALLET/cmd/utils"
	"github.com/ofbank_wallet/OFBANK_WALLET/node"
	"github.com/ofbank_wallet/OFBANK_WALLET/console"
)

var(
	consoleFlags = []cli.Flag{utils.JSpathFlag, utils.ExecFlag, utils.PreloadJSFlag}

	consoleCommand = cli.Command{
		Action:   utils.MigrateFlags(localConsole),
		Name:     "console",
		Usage:    "Start an interactive JavaScript environment",
		Flags:    /*append(*/append(nodeFlags, consoleFlags...)/*, whisperFlags...)*/,
		Category: "CONSOLE COMMANDS",
		Description: `
The Geth console is an interactive shell for the JavaScript runtime environment
which exposes a node admin interface as well as the Ðapp JavaScript API.
See https://github.com/ethereum/go-ethereum/wiki/Javascipt-Console.`,
	}
)

func localConsole(ctx *cli.Context) error {
	newNode,err:=node.NewNode()

	if err != nil {
		utils.Fatalf("Failed to attach to the node : %v", err)
		return nil
	}
	client,err:=newNode.Attach()

	if err != nil {
		utils.Fatalf("Failed to attach to the inproc geth: %v", err)
		return nil
	}
	config := console.Config{
		DataDir: utils.MakeDataDir(ctx),
		DocRoot: ctx.GlobalString(utils.JSpathFlag.Name),
		Client:  client,
		Preload: utils.MakeConsolePreloads(ctx),
	}

	console, err := console.New(config)
	if err != nil {
		utils.Fatalf("Failed to start the JavaScript console: %v", err)
	}

	defer console.Stop(false)
	// If only a short execution was requested, evaluate and return
	if script := ctx.GlobalString(utils.ExecFlag.Name); script != "" {
		console.Evaluate(script)
		return nil
	}
	// Otherwise print the welcome screen and enter interactive mode
	console.Welcome()
	console.Interactive()

	return nil
}