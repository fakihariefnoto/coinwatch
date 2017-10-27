package notifier

import (
	"log"
	"os/exec"
)

type (
	Message struct {
		Title string
		Body  string
		Icon  int
	}

	INotifier interface {
		NewConfig(Config) Notifier
	}
	pkgNotify struct{}

	Notifier interface {
		Message(string, string) Sender
	}
	Config struct {
		FixedTitle string
		Icon       int
	}

	Sender interface {
		Send()
		SendWithIcon(int)
	}
)

func New() INotifier {
	return &pkgNotify{}
}

func (p *pkgNotify) NewConfig(conf Config) Notifier {
	return &conf
}

func (c *Config) Message(title, body string) Sender {

	message := Message{
		Title: title,
		Body:  body,
	}

	if c.FixedTitle != "" {
		message.Title = c.FixedTitle
	}
	return message
}

func (s Message) Send() {
	if s.Icon == 0 {
		s.Icon = SMART
	}
	executionNotify(s.Icon, s.Title, s.Body)
}

func (s Message) SendWithIcon(icon int) {
	executionNotify(icon, s.Title, s.Body)
}

func executionNotify(icon int, title, body string) {
	getIcon := messageIcon[icon]
	exe := exec.Command("notify-send", "-i", getIcon, title, body)
	err := exe.Run()
	if err != nil {
		log.Println("upps got error ->", err)
	}
}
