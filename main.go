package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Migan178/HennaDM/commands"
	"github.com/Migan178/HennaDM/configs"
	"github.com/bwmarrin/discordgo"
)

func main() {
	err := dg.Open()
	if err != nil {
		panic(err)
	}

	defer dg.Close()

	var globalCmds []*discordgo.ApplicationCommand
	var developerOnlyGuildCmds []*discordgo.ApplicationCommand
	for _, cmd := range commands.GetDiscommand().Commands {
		if cmd.Flags&commands.CommandFlagsIsDeveloperOnlyCommand != 0 {
			developerOnlyGuildCmds = append(developerOnlyGuildCmds, cmd.ApplicationCommand)
			continue
		}

		globalCmds = append(globalCmds, cmd.ApplicationCommand)
	}

	_, err = dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, "", globalCmds)
	if err != nil {
		fmt.Println(err)
	}

	if len(developerOnlyGuildCmds) != 0 && configs.GetConfig().Command.DeveloperOnlyGuildID != "" {
		_, err = dg.ApplicationCommandBulkOverwrite(dg.State.User.ID, configs.GetConfig().Command.DeveloperOnlyGuildID, developerOnlyGuildCmds)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("[MF] The bot is running. version:", configs.MFVersion)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
