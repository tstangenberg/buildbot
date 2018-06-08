package buildbot

import (
	"gopkg.in/urfave/cli.v2"
	"fmt"
)

func NewCli() *cli.App {
	app := &cli.App{Name: "Build Bot",
		Usage: "#1 Software Development Automation Robot",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			bot := factory.NewBot()
			bot.HelpTheDeveloper()
			return nil
		},
	}
	return app;
}
