package handler

import (
	"fmt"

	"github.com/Migan178/HennaDM/commands"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if err := commands.GetDiscommand().ChatInputRun(i.ApplicationCommandData().Name, s, i); err != nil {
				fmt.Println(err)
			}
		case discordgo.InteractionMessageComponent:
			if err := commands.GetDiscommand().ComponentRun(s, i); err != nil {
				fmt.Println(err)
			}
		case discordgo.InteractionModalSubmit:
			if err := commands.GetDiscommand().ModalRun(s, i); err != nil {
				fmt.Println(err)
			}
		case discordgo.InteractionApplicationCommandAutocomplete:
			if err := commands.GetDiscommand().ChatInputAutocomplete(i.ApplicationCommandData().Name, s, i); err != nil {
				fmt.Println(err)
			}
		}
	}()
}
