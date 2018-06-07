package buildbot

import "gopkg.in/urfave/cli.v2"

func NewCli() *cli.App {
	app := &cli.App{Name: "Build Bot",
		Usage: "#1 Software Development Automation Robot",
	}
	return app;
}
