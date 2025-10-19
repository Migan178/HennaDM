package main

import (
	_ "github.com/Muffin-laboratory/mf/commands"
	_ "github.com/Muffin-laboratory/mf/components"
	"github.com/Muffin-laboratory/mf/configs"
	"github.com/Muffin-laboratory/mf/handler"
	_ "github.com/Muffin-laboratory/mf/modals"
	"github.com/bwmarrin/discordgo"
)

var dg *discordgo.Session

func init() {
	dg, _ = discordgo.New("Bot " + configs.GetConfig().Bot.Token)

	// Handler
	go dg.AddHandler(handler.InteractionCreate)
}
