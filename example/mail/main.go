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
		from:    "Jane Doe <jane@doe.com>",
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
	mailHeaders.SetSizePolicy(tui.Expanding, tui.Minimum)
	mailHeaders.AppendRow(tui.NewLabel("From:"), fromLabel)
	mailHeaders.AppendRow(tui.NewLabel("Subject:"), subjLabel)
	mailHeaders.AppendRow(tui.NewLabel("Date:"), dateLabel)

	// Panel for reading the mail contents.
	mailView := tui.NewVBox(
		mailHeaders,
		bodyLabel,
	)
	mailView.SetSizePolicy(tui.Expanding, tui.Expanding)

	// List with all the mail items.
	mailList := tui.NewTable(0, 0)
	mailList.SetSizePolicy(tui.Expanding, tui.Minimum)

	for _, m := range mails {
		mailList.AppendRow(
			tui.NewLabel(m.subject),
			tui.NewLabel(m.from),
			tui.NewLabel(m.date),
		)
	}

	mailList.SetColumnStretch(0, 3)
	mailList.SetColumnStretch(1, 2)
	mailList.SetColumnStretch(2, 1)

	//mailList.SetSizePolicy(tui.Expanding, tui.Expanding)
	mailList.OnSelectionChanged(func(t *tui.Table) {
		m := mails[t.Selected()]
		fromLabel.SetText(m.from)
		subjLabel.SetText(m.subject)
		dateLabel.SetText(m.date)
		bodyLabel.SetText(m.body)
	})

	// Main layout for the application.
	root := tui.NewVBox(
		mailList,
		tui.NewLabel(""),
		mailView,
	)
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
