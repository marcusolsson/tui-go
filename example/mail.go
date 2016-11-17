package main

import (
	"fmt"

	"github.com/marcusolsson/tui-go"
)

type mail struct {
	from    string
	subject string
	date    string
	body    string
}

func (m mail) String() string {
	return fmt.Sprintf("%s  %s  %s", m.subject, m.from, m.date)
}

var mails = []mail{
	{
		from:    "John Doe <john@doe.com>",
		subject: "Vacation pictures",
		date:    "Yesterday",
		body: `
Hey,

Where can I find the pictures from the diving trip?

Cheers,
John`,
	},
	{
		from:    "Jane Doe <john@doe.com>",
		subject: "Meeting notes",
		date:    "Yesterday",
		body: `
Here are the notes from today's meeting.

/Jane`,
	},
}

func main() {
	var (
		fromLabel = tui.NewLabel("")
		subjLabel = tui.NewLabel("")
		dateLabel = tui.NewLabel("")
		bodyLabel = tui.NewLabel("")
	)

	mailHeaders := tui.NewGrid(0, 0)
	mailHeaders.AppendRow(tui.NewLabel("From:"), fromLabel)
	mailHeaders.AppendRow(tui.NewLabel("Subject:"), subjLabel)
	mailHeaders.AppendRow(tui.NewLabel("Date:"), dateLabel)

	// Panel for reading the mail contents.
	mailView := tui.NewVerticalBox(
		mailHeaders,
		bodyLabel,
	)
	mailView.SetBorder(true)
	mailView.SetSizePolicy(tui.Expanding, tui.Expanding)

	// List with all the mail items.
	mailList := tui.NewList()
	for _, m := range mails {
		mailList.AddItems(m.String())
	}
	mailList.SetSizePolicy(tui.Expanding, tui.Expanding)
	mailList.SetRows(10)
	mailList.SetSelected(0)
	mailList.OnSelectionChanged(func(l *tui.List) {
		m := mails[l.Selected()]
		fromLabel.SetText(m.from)
		subjLabel.SetText(m.subject)
		dateLabel.SetText(m.date)
		bodyLabel.SetText(m.body)
	})

	// Panel containing all the mails.
	inboxView := tui.NewVerticalBox(mailList)
	inboxView.SetBorder(true)
	inboxView.SetSizePolicy(tui.Expanding, tui.Expanding)

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
