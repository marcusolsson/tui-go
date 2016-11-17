package main

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
)

type mail struct {
	from    string
	subject string
	date    string
}

func (m mail) String() string {
	return fmt.Sprintf("%s  %s  %s", m.subject, m.from, m.date)
}

var mails = []mail{
	{from: "John Doe <john@doe.com>", subject: "Vacation pictures", date: "Yesterday"},
	{from: "Jane Doe <john@doe.com>", subject: "Meeting notes", date: "Yesterday"},
}

func main() {

	// List with all the mail items.
	mailList := tui.NewList()
	for _, m := range mails {
		mailList.AddItems(m.String())
	}
	mailList.SetRows(10)
	mailList.SetSelected(0)
	mailList.OnItemActivated(func(l *tui.List) {

	})

	// Panel containing all the mails.
	inboxView := tui.NewVerticalBox(mailList)
	inboxView.SetBorder(true)
	inboxView.SetSizePolicy(tui.Expanding, tui.Expanding)

	// Panel for reading the mail contents.
	mailView := tui.NewVerticalBox(
		tui.NewLabel("From: Me"),
		tui.NewLabel("Subject: Vacation pictures"),
	)
	mailView.SetBorder(true)
	mailView.SetSizePolicy(tui.Expanding, tui.Expanding)

	// Main layout for the application.
	root := tui.NewVerticalBox(inboxView, mailView)
	root.SetSizePolicy(tui.Expanding, tui.Expanding)

	// Start the application.
	ui := tui.New(root)
	ui.SetShortcut('h', func() {
		if mailView.IsVisible() {
			mailView.Hide()
		} else {
			mailView.Show()
		}
	})
	if err := ui.Run(); err != nil {
		panic(err)
	}
}
