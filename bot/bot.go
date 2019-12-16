package bot

import (
	"fmt"
	"strings"
	"time"

	"github.com/shvgn/work-schedule-bot/table"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	// User name substring
	name = "Shevchenko"
)

// InitBot inits all bot logic
func InitBot(token string, tbl *table.Table) (*tb.Bot, error) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 15 * time.Second},
	})

	if err != nil {
		return nil, err
	}

	// This button will be displayed in user's reply keyboard.
	setBtn := tb.ReplyButton{Text: "Add time"}
	clearBtn := tb.ReplyButton{Text: "Clear"}

	replyKeys := [][]tb.ReplyButton{
		{setBtn},
		{clearBtn},
	}

	replyOpts := &tb.ReplyMarkup{
		ReplyKeyboard: replyKeys,
	}

	b.Handle(&setBtn, func(m *tb.Message) {
		rec, err := tbl.Append(name, m.Time())
		var text string
		if err != nil {
			text = fmt.Sprintf("😨 Oops, %v", err)
		} else {
			text = formatCurrentTime(rec.Starts, rec.Ends)
		}
		b.Send(m.Sender, text, replyOpts, tb.ModeMarkdown)
	})

	b.Handle(&clearBtn, func(m *tb.Message) {
		tbl.Clear(name, m.Time())
		var text string
		if err != nil {
			text = fmt.Sprintf("😨 Oops, %v", err)
		} else {
			text = "All cleared!"
		}
		b.Send(m.Sender, text, replyOpts)
	})

	// Command: /start <PAYLOAD>
	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		if m.Payload != "dear settime" {
			return
		}

		text := "Hello"
		b.Send(m.Sender, text, replyOpts)
	})

	return b, nil
}

func formatCurrentTime(start, end []string) string {
	b := strings.Builder{}
	b.WriteString("```\n")

	for i := range end {
		b.WriteString(start[i])
		b.WriteString("  ")
		b.WriteString(end[i])
		b.WriteString("\n")
	}

	if len(start) > len(end) {
		b.WriteString(start[len(start)-1])
	}
	b.WriteString("\n```")
	return b.String()
}
