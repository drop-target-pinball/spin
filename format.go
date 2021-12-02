package spin

import (
	"fmt"
	"reflect"
	"strings"

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

func formatActionOrEvent(a interface{}) string {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)
	fields := make([]string, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fields = append(fields, fmt.Sprintf("%v=%v", f.Name, v.Field(i)))
	}
	return t.Name() + " " + strings.Join(fields, ", ")
}

func FormatAction(a Action) string {
	return formatActionOrEvent(a)
}

func FormatEvent(e Event) string {
	return formatActionOrEvent(e)
}
