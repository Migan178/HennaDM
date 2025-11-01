package handler

import (
	"fmt"
	"strings"

	"github.com/Migan178/HennaDM/configs"
	"github.com/bwmarrin/discordgo"
)

func GuildMemberRemove(s *discordgo.Session, m *discordgo.GuildMemberRemove) {
	go func() {
		channels, err := s.GuildChannels(m.GuildID)
		if err != nil {
			fmt.Println(err)
			return
		}

		var hennaDMClosedCategory, userChannel *discordgo.Channel

		for _, c := range channels {
			if c.Type == discordgo.ChannelTypeGuildCategory && c.Name == configs.GetConfig().Bot.ClosedCategoryName {
				hennaDMClosedCategory = c

				continue
			}

			if strings.Contains(c.Topic, m.User.ID) && c.ParentID != hennaDMClosedCategory.ID {
				userChannel = c

				continue
			}

		}

		if _, err := s.ChannelEditComplex(userChannel.ID, &discordgo.ChannelEdit{
			ParentID: hennaDMClosedCategory.ID,
			PermissionOverwrites: []*discordgo.PermissionOverwrite{
				{
					ID:   m.GuildID,
					Type: discordgo.PermissionOverwriteTypeRole,
					Deny: discordgo.PermissionAll,
				},
			},
		}); err != nil {
			fmt.Println(err)
			return
		}
	}()
}
