package buildbot

//go:generate moq -out factory_moq_test.go . Factory

type Factory interface {
	NewBot() Bot
	NewConfiguration() Configuration
}

type DefaultFactory struct {
}

func (f *DefaultFactory) NewBot() Bot  {
	bot := Bot{}
	configuration := factory.NewConfiguration()
	configuration.Initialize()
	bot.config = configuration
	return bot
}

func (f *DefaultFactory) NewConfiguration() Configuration  {
	c := Configuration{}
	return c
}

var factory Factory

func init() {
	factory = &DefaultFactory{}
}


