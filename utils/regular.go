package utils

import "regexp"

func IsMobile(m string) bool {
	regRuler := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRuler)

	return reg.MatchString(m)
}
