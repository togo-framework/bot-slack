# bot-slack — docs

**Slack bot.** Slack driver via slack-go Socket Mode — events/slash commands → the registry.

## Install

```bash
togo install togo-framework/bot-slack
```

Registers on the [`bot`](https://github.com/togo-framework/bot) base; select it with **BOT_DRIVER (or bot.provider)**, then use **`togo bot`**.

## Interface

`Bot` — `Start`/`Stop`/`Send`, plus a command/handler registry (`OnCommand`/`OnMessage`) so any plugin can add bot commands.

## Configuration

| Env var | Description |
|---|---|
| `SLACK_BOT_TOKEN` | Slack bot token `xoxb-…` (required). |
| `SLACK_APP_TOKEN` | Slack app-level token `xapp-…` for Socket Mode (required). |

## Usage & notes

Uses Socket Mode (needs both `xoxb-` bot and `xapp-` app tokens); posts messages/blocks.

## Example

```bash
togo bot:send '#general' 'Deployed!'
togo bot:ask 'summarize the latest release'
```

## Links

- [slack-go](https://github.com/slack-go/slack)
- [Marketplace](https://to-go.dev/marketplace)
- [Source](https://github.com/togo-framework/bot-slack)
