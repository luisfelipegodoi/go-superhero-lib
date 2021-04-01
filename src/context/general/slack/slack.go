package slack

import (
	"github.com/multiplay/go-slack/chat"
	"github.com/multiplay/go-slack/lrhook"
	"github.com/sirupsen/logrus"
)

type Slack struct {
	LogLevel       logrus.Level
	SlackChannel   string
	SlackIconEmoji string
	WebHookUrl     string
}

func NewSlack(logLevel logrus.Level, slackChannel, slackIconEmoji, webHookUrl string) Slack {
	return Slack{
		LogLevel:       logLevel,
		SlackChannel:   slackChannel,
		SlackIconEmoji: slackIconEmoji,
		WebHookUrl:     webHookUrl,
	}
}

func (sa *Slack) SlackSendAlerts() *lrhook.Hook {
	cfg := lrhook.Config{
		MinLevel: sa.LogLevel,
		Message: chat.Message{
			Channel:   sa.SlackChannel,
			IconEmoji: sa.SlackIconEmoji,
		},
	}

	webHook := lrhook.New(cfg, sa.WebHookUrl)
	return webHook
}
