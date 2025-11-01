package handler

import (
	"fmt"

	"github.com/Migan178/HennaDM/configs"
	"github.com/bwmarrin/discordgo"
)

func GuildMemberAdd(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
	go func() {
		channels, err := s.GuildChannels(m.GuildID)
		if err != nil {
			fmt.Println(err)
			return
		}

		var hennaDMOpenCategory *discordgo.Channel

		for _, c := range channels {
			if c.Type == discordgo.ChannelTypeGuildCategory && c.Name == configs.GetConfig().Bot.OpenCategoryName {
				hennaDMOpenCategory = c
				break
			}
		}

		if _, err = s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
			Type:     discordgo.ChannelTypeGuildText,
			Name:     m.User.Username,
			ParentID: hennaDMOpenCategory.ID,
			Topic:    m.User.ID,
			PermissionOverwrites: []*discordgo.PermissionOverwrite{
				{
					Type:  discordgo.PermissionOverwriteTypeMember,
					ID:    m.User.ID,
					Allow: discordgo.PermissionAllText,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return
		}
	}()

}
