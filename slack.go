// Package slack is the Slack driver for togo's bot subsystem, using Socket Mode
// (no public webhook URL required). Blank-import it alongside
// github.com/togo-framework/bot and set BOT_DRIVER=slack plus SLACK_APP_TOKEN
// (xapp-…) and SLACK_BOT_TOKEN (xoxb-…).
//
//	import _ "github.com/togo-framework/bot"
//	import _ "github.com/togo-framework/bot-slack"
package slack

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"github.com/togo-framework/bot"
	"github.com/togo-framework/togo"
)

func init() {
	bot.RegisterDriver("slack", makeDriver)
}

func makeDriver(k *togo.Kernel, dispatch func(context.Context, bot.Message)) (bot.Bot, error) {
	appToken := os.Getenv("SLACK_APP_TOKEN")
	botToken := os.Getenv("SLACK_BOT_TOKEN")
	if appToken == "" || botToken == "" {
		return nil, fmt.Errorf("bot-slack: SLACK_APP_TOKEN (xapp-…) and SLACK_BOT_TOKEN (xoxb-…) must both be set")
	}
	if !strings.HasPrefix(appToken, "xapp-") {
		return nil, fmt.Errorf("bot-slack: SLACK_APP_TOKEN must start with xapp- (Socket Mode app-level token)")
	}
	api := slack.New(botToken, slack.OptionAppLevelToken(appToken))
	client := socketmode.New(api)
	return &driver{api: api, client: client, dispatch: dispatch}, nil
}

type driver struct {
	api      *slack.Client
	client   *socketmode.Client
	dispatch func(context.Context, bot.Message)
	botID    string
}

// Start runs the Socket Mode event loop, dispatching message events. It blocks
// until ctx is canceled.
func (d *driver) Start(ctx context.Context) error {
	// Resolve our own user ID so we can ignore our own messages.
	if auth, err := d.api.AuthTestContext(ctx); err == nil {
		d.botID = auth.UserID
	}

	go func() {
		for evt := range d.client.Events {
			if evt.Type != socketmode.EventTypeEventsAPI {
				continue
			}
			eventsAPI, ok := evt.Data.(slackevents.EventsAPIEvent)
			if !ok {
				continue
			}
			d.client.Ack(*evt.Request)
			if eventsAPI.Type != slackevents.CallbackEvent {
				continue
			}
			if mev, ok := eventsAPI.InnerEvent.Data.(*slackevents.MessageEvent); ok {
				// Skip bot messages, edits, and our own messages.
				if mev.BotID != "" || mev.SubType != "" || mev.User == "" || mev.User == d.botID {
					continue
				}
				d.dispatch(ctx, bot.Message{
					Channel:  mev.Channel,
					User:     mev.User,
					Username: mev.User,
					Text:     mev.Text,
					Platform: "slack",
					Raw:      map[string]any{"event": mev},
				})
			}
		}
	}()

	go func() {
		<-ctx.Done()
	}()
	// RunContext blocks until ctx is canceled or the connection drops.
	return d.client.RunContext(ctx)
}

// Stop is a no-op beyond context cancellation; the socketmode loop ends when the
// context passed to Start is canceled.
func (d *driver) Stop() error { return nil }

// Send posts msg to a channel ID using the Web API.
func (d *driver) Send(ctx context.Context, channel, msg string) error {
	_, _, err := d.api.PostMessageContext(ctx, channel, slack.MsgOptionText(msg, false))
	if err != nil {
		return fmt.Errorf("bot-slack: send: %w", err)
	}
	return nil
}
