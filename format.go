package spin

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var locale = message.NewPrinter(language.English)

func FormatScore(format string, a ...interface{}) string {
	score := locale.Sprintf(format, a...)
	if score == "0" {
		return "00"
	}
	return score
}
