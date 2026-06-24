---
name: bot-slack
description: Run a Slack bot in a togo app (Socket Mode) — configure BOT_DRIVER=slack + SLACK_APP_TOKEN/SLACK_BOT_TOKEN and register handlers with the bot plugin
---

# togo bot-slack

Slack driver for the togo `bot` subsystem, using **Socket Mode** (no public URL).

## Setup

```bash
togo install togo-framework/bot
togo install togo-framework/bot-slack
```

1. Create an app at [api.slack.com/apps](https://api.slack.com/apps).
2. **Enable Socket Mode**, generate an App-Level Token (`xapp-…`) with
   `connections:write`.
3. Add bot scopes (`chat:write`, `channels:history`, `im:history`,
   `app_mentions:read`), install the app → Bot Token (`xoxb-…`).
4. Subscribe to `message.channels` / `message.im` events.
5. `.env`:
   ```bash
   BOT_DRIVER=slack
   SLACK_APP_TOKEN=xapp-...
   SLACK_BOT_TOKEN=xoxb-...
   ```
6. Register handlers with `bot.OnCommand` / `bot.OnMessage` (see the `bot` skill).

## Notes
- `m.Channel` is the Slack channel ID; reply with `Service.Send`.
- Socket Mode means it runs fine from localhost/behind NAT.
- The driver ignores bot/edited messages (no loops).
- Never commit tokens; keep them in `.env`.
