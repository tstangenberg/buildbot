package buildbot

import "github.com/labstack/gommon/log"

type Bot struct {
	config Configuration
}

func (b *Bot) HelpTheDeveloper() {
	task := b.config.Task
	log.Info("running")
	log.Info(task.Command)
	log.Info(task.Arguments)
}
