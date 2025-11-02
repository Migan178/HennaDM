package commands

import (
	"github.com/Migan178/HennaDM/builders"
	"github.com/Migan178/HennaDM/configs"
	"github.com/bwmarrin/discordgo"
)

var SetCommand = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "설정",
		Description: "해당 길드에 봇에게 필요한 설정을 해요.",
	},
	Run: func(inter *builders.InteractionCreate) error {
		var existOpenCategory, existClosedCategory bool

		channels, err := inter.Session.GuildChannels(inter.GuildID)
		if err != nil {
			return err
		}

		for _, ch := range channels {
			if ch.Name == configs.GetConfig().Bot.OpenCategoryName && ch.Type == discordgo.ChannelTypeGuildCategory {
				existOpenCategory = true

				continue
			}

			if ch.Name == configs.GetConfig().Bot.ClosedCategoryName && ch.Type == discordgo.ChannelTypeGuildCategory {
				existClosedCategory = true

				continue
			}
		}

		if existOpenCategory && existClosedCategory {
			return builders.NewMessageSender(inter).
				AddComponents(builders.MakeErrorContainer("이미 해당 길드는 설정이 돼 있어요.")).
				SetComponentsV2(true).
				SetEphemeral(true).
				Send()
		}

		permissionOverwrites := []*discordgo.PermissionOverwrite{
			{
				ID:   inter.GuildID,
				Type: discordgo.PermissionOverwriteTypeRole,
				Deny: discordgo.PermissionAll,
			},
			{
				ID:    inter.Session.State.User.ID,
				Type:  discordgo.PermissionOverwriteTypeMember,
				Allow: discordgo.PermissionAllChannel,
			},
		}

		if !existOpenCategory {
			if _, err = inter.Session.GuildChannelCreateComplex(inter.GuildID, discordgo.GuildChannelCreateData{
				Name:                 "hdm-open",
				Type:                 discordgo.ChannelTypeGuildCategory,
				PermissionOverwrites: permissionOverwrites,
			}); err != nil {
				return err
			}
		}

		if !existClosedCategory {
			if _, err := inter.Session.GuildChannelCreateComplex(inter.GuildID, discordgo.GuildChannelCreateData{
				Name:                 "hdm-closed",
				Type:                 discordgo.ChannelTypeGuildCategory,
				PermissionOverwrites: permissionOverwrites,
			}); err != nil {
				return err
			}
		}

		return builders.NewMessageSender(inter).
			AddComponents(builders.MakeSuccessContainer("해당 길드에 설정을 완료했어요.")).
			SetComponentsV2(true).
			SetEphemeral(true).
			Send()
	},
}

func init() {
	GetDiscommand().LoadCommand(SetCommand)
}
