package utils

import (
	"strconv"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}

func ToCurrencyStr(value float64) string {
	p := message.NewPrinter(language.Portuguese)
	return p.Sprintf("%.2f", value)
}
