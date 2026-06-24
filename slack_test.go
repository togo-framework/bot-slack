package slack

import (
	"context"
	"os"
	"testing"

	"github.com/togo-framework/bot"
)

func TestDriverRegistered(t *testing.T) {
	found := false
	for _, n := range bot.Drivers() {
		if n == "slack" {
			found = true
		}
	}
	if !found {
		t.Fatal("slack driver not registered with bot")
	}
}

func TestFactoryRequiresTokens(t *testing.T) {
	oldApp, oldBot := os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_BOT_TOKEN")
	_ = os.Unsetenv("SLACK_APP_TOKEN")
	_ = os.Unsetenv("SLACK_BOT_TOKEN")
	defer func() {
		os.Setenv("SLACK_APP_TOKEN", oldApp)
		os.Setenv("SLACK_BOT_TOKEN", oldBot)
	}()

	if b, err := makeDriver(nil, nil); err == nil || b != nil {
		t.Fatalf("expected error without tokens, got (%v,%v)", b, err)
	}
}

func TestFactoryRejectsBadAppToken(t *testing.T) {
	oldApp, oldBot := os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_BOT_TOKEN")
	os.Setenv("SLACK_APP_TOKEN", "not-an-app-token")
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-fake")
	defer func() {
		os.Setenv("SLACK_APP_TOKEN", oldApp)
		os.Setenv("SLACK_BOT_TOKEN", oldBot)
	}()

	if b, err := makeDriver(nil, nil); err == nil || b != nil {
		t.Fatalf("expected error for non-xapp app token, got (%v,%v)", b, err)
	}
}

func TestFactoryBuildsClient(t *testing.T) {
	oldApp, oldBot := os.Getenv("SLACK_APP_TOKEN"), os.Getenv("SLACK_BOT_TOKEN")
	os.Setenv("SLACK_APP_TOKEN", "xapp-fake-token")
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-fake-token")
	defer func() {
		os.Setenv("SLACK_APP_TOKEN", oldApp)
		os.Setenv("SLACK_BOT_TOKEN", oldBot)
	}()

	b, err := makeDriver(nil, func(_ context.Context, _ bot.Message) {})
	if err != nil || b == nil {
		t.Fatalf("expected a driver with valid-shaped tokens, got (%v,%v)", b, err)
	}
}
