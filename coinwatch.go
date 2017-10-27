package main

import (
	"fmt"
	notif "github.com/fakihariefnoto/coinwatch/notifier"
	"time"
)

var (
	notify notif.INotifier
	sender notif.Notifier
)

func init() {
	notify = notif.New()
	sender = notify.NewConfig(notif.Config{
		Icon:       notif.ANGRY,
		FixedTitle: "Percobaan",
	})
}

func main() {
	fmt.Println("Trial send notification")
	sender.Message("", "wkwkwkwk").Send()

	time.Sleep(1000)

	sender.Message("kirim", "hahahaha").SendWithIcon(notif.ANGEL)

	time.Sleep(1000)

	sender.Message("hmmmmm", "hmmmmmm").SendWithIcon(notif.LAUGH)
}
