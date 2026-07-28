package service

import (
	"strings"

	"github.com/luzhkovn/go1fl-sprint6-final-tpl/pkg/morse"
)

func AutoConvert(input string) (string, error) {
	MorseOrText := strings.Map(func(r rune) rune {
		if r == ' ' || r == '.' || r == '-' {
			return -1
		}

		return r
	}, input)

	if MorseOrText == "" {
		MorseToText := morse.ToText(input)
		return MorseToText, nil
	} else {
		TextInMorse := morse.ToMorse(input)
		return TextInMorse, nil

	}

}
