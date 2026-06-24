<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/bot-slack</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/bot-slack"><img src="https://pkg.go.dev/badge/github.com/togo-framework/bot-slack.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Slack driver for the <a href="https://github.com/togo-framework/bot">togo bot</a> subsystem.</strong></p>
</div>

## Install

```bash
togo install togo-framework/bot
togo install togo-framework/bot-slack
```

<!-- /togo-header -->

The **Slack** driver for togo's [`bot`](https://github.com/togo-framework/bot)
subsystem, built on [slack-go](https://github.com/slack-go/slack) in **Socket
Mode** — no public webhook URL required, so it works from localhost and behind
NAT. It dispatches each message to the bot command/message handlers you register
once with `bot.OnCommand` / `bot.OnMessage`.

## Configure

1. Create a Slack app at [api.slack.com/apps](https://api.slack.com/apps).
2. Enable **Socket Mode** and generate an **App-Level Token** (`xapp-…`) with the
   `connections:write` scope.
3. Add bot scopes (`app_mentions:read`, `channels:history`, `chat:write`,
   `im:history`) and install the app to your workspace to get the **Bot Token**
   (`xoxb-…`).
4. Subscribe to the `message.channels` / `message.im` events.
5. Set env:
   ```bash
   BOT_DRIVER=slack
   SLACK_APP_TOKEN=xapp-...
   SLACK_BOT_TOKEN=xoxb-...
   ```

Blank-import the driver next to the base:

```go
import (
	_ "github.com/togo-framework/bot"
	_ "github.com/togo-framework/bot-slack"
)
```

`m.Channel` is the Slack channel ID; `Service.Send(ctx, channel, msg)` posts to
it. The driver ignores bot messages and edits to avoid loops.

## License

MIT © togo-framework

<!-- togo-sponsors -->
---

<div align="center">
  <h3>Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><strong>ID8 Media</strong></a> &nbsp;·&nbsp;
    <a href="https://one-studio.co"><strong>One Studio</strong></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
