package main

import (
	_ "github.com/Migan178/HennaDM/commands"
	_ "github.com/Migan178/HennaDM/components"
	"github.com/Migan178/HennaDM/configs"
	"github.com/Migan178/HennaDM/handler"
	_ "github.com/Migan178/HennaDM/modals"
	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func init() {
	dg, _ = discordgo.New("Bot " + configs.GetConfig().Bot.Token)

	dg.Identify.Intents = discordgo.IntentGuildMembers | discordgo.IntentsGuilds

	// Handler
	go dg.AddHandler(handler.InteractionCreate)
	go dg.AddHandler(handler.GuildMemberAdd)
}
