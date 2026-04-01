package theInterface

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func HexToDec(hexStr string) (int64, error) {
	dec, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hex string: %s", hexStr)
	}
	return dec, nil
}

func ReplaceHexWithDec(input string) string {
	re := regexp.MustCompile(`(?i)\b([0-9a-f]+) \(hex\)`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		parts := strings.Split(match, " ")
		hex := parts[0]
		dec, err := HexToDec(hex)
		if err != nil {
			return match
		}
		return fmt.Sprintf("%d", dec)
	})
}