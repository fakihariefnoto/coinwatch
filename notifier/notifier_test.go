package notifier

import (
	"testing"
)

var (
	notify INotifier
	sender Notifier
)

func TestAll(t *testing.T) {
	notify = New()
	sender = notify.NewConfig(Config{
		Icon:       ANGRY,
		FixedTitle: "Percobaan",
	})

	sender.Message("", "wkwkwkwk").Send()

	sender.Message("kirim", "hahahaha").SendWithIcon(ANGEL)

	sender.Message("hmmmmm", "hmmmmmm").SendWithIcon(LAUGH)
}
