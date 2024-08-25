package config

import "flag"

type CommandArgs struct {
	Config string
}

var Args *CommandArgs

func newCommandArgs() {
	config := flag.String("config", "config.yaml", "config file")
	flag.Parse()
	Args = &CommandArgs{}
	Args.Config = *config
}
