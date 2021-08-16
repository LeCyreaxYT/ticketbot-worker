package setup

import (
	"github.com/TicketsBot/common/permission"
	"github.com/TicketsBot/worker/bot/command"
	"github.com/TicketsBot/worker/bot/command/registry"
	"github.com/TicketsBot/worker/bot/dbclient"
	"github.com/TicketsBot/worker/bot/i18n"
	"github.com/TicketsBot/worker/bot/utils"
	"github.com/rxdn/gdl/objects/interaction"
	"github.com/rxdn/gdl/rest/request"
)

type TranscriptsSetupCommand struct{}

func (TranscriptsSetupCommand) Properties() registry.Properties {
	return registry.Properties{
		Name:            "transcripts",
		Description:     i18n.HelpSetup,
		Aliases:         []string{"transcript", "archives", "archive"},
		PermissionLevel: permission.Admin,
		Category:        command.Settings,
		Arguments: command.Arguments(
			command.NewRequiredArgument("channel", "The channel that ticket transcripts should be sent to", interaction.OptionTypeChannel, i18n.SetupTranscriptsInvalid),
		),
	}
}

func (c TranscriptsSetupCommand) GetExecutor() interface{} {
	return c.Execute
}

func (TranscriptsSetupCommand) Execute(ctx registry.CommandContext, channelId uint64) {
	if _, err := ctx.Worker().GetChannel(channelId); err != nil {
		if restError, ok := err.(request.RestError); ok && restError.IsClientError() {
			ctx.Reply(utils.Red, "Error", i18n.SetupTranscriptsInvalid, ctx.ChannelId)
			ctx.Reject()
		} else {
			ctx.HandleError(err)
		}

		return
	}

	if err := dbclient.Client.ArchiveChannel.Set(ctx.GuildId(), channelId); err == nil {
		ctx.Accept()
		ctx.Reply(utils.Green, "Setup", i18n.SetupTranscriptsComplete, channelId)
	} else {
		ctx.HandleError(err)
	}
}
