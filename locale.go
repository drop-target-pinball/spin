package spin

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var locale = message.NewPrinter(language.English)

func Sprintf(format string, a ...interface{}) string {
	return locale.Sprintf(format, a...)
}
