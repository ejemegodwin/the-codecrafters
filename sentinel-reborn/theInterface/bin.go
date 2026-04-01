package theInterface

import (
	"strconv"
	"strings"
)

func BinToDec(binstr string) string {

	result := []string{}
	s := strings.Fields(binstr)
	for i := 0; i < len(s); i++ {
		if s[i] == "(bin)" {
			dec, err := strconv.ParseInt(result[len(result)-1], 2, 64)
			if err != nil {
				return ""
			}
			result[len(result)-1] = strconv.FormatInt(dec, 10)

			continue

		}
		result = append(result, s[i])
	}
	return strings.Join(result, " ")

}
