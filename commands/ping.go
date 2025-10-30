package commands

import (
	"fmt"

	"github.com/Migan178/HennaDM/builders"
	"github.com/bwmarrin/discordgo"
)

var PingCommand = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "핑",
		Description: "봇의 레이턴시를 확인해요.",
	},
	Run: func(inter *builders.InteractionCreate) error {
		title := fmt.Sprintf("### 🏓 %s의 지연시간", inter.Session.State.User.Username)

		if err := builders.NewMessageSender(inter).
			AddComponents(builders.ContainerBuilder().AddText(title).AddText("- 지연시간 측정 중...")).
			SetComponentsV2(true).
			Send(); err != nil {
			return err
		}

		message, err := inter.FetchReply()
		if err != nil {
			return err
		}

		createdTimestamp, _ := discordgo.SnowflakeTimestamp(inter.ID)
		discordPing := message.Timestamp.Sub(createdTimestamp).Milliseconds()

		return builders.NewMessageSender(inter).
			AddComponents(
				builders.ContainerBuilder().
					AddText(title).
					AddText(fmt.Sprintf("- **디스코드 지연시간:** `%d`ms", discordPing)),
			).
			SetComponentsV2(true).
			Send()
	},
}

func init() {
	GetDiscommand().LoadCommand(PingCommand)
}
