package bytefmt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	b  = int64(1)
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
	tb = 1024 * gb
)

// Convert raw string 1B, 1K, 1Kb ... to byte
func HumanToByte(rawString string) (int64, error) {
	var bytes int64
	pattern := regexp.MustCompile(`^([0-9]+)([KMGT])?(B|b)?$`)
	tokens := pattern.FindStringSubmatch(rawString)
	// If we passed correct raw string,
	// FindStringSubmatch return slice with 4 elements
	if len(tokens) != 4 {
		return 0, errors.New("Incorrect string value for conversion")
	}
	value, err := strconv.ParseInt(tokens[1], 10, 64)
	if err != nil {
		return value, err
	}

	switch strings.ToUpper(tokens[2]) {
	case "B":
		bytes = value * b
	case "K":
		bytes = value * kb
	case "M":
		bytes = value * mb
	case "G":
		bytes = value * gb
	case "T":
		bytes = value * tb
	default:
		bytes = value * b
	}

	return bytes, nil
}

// Convert byte to human representation value, like that
// 1024 - 1K or 1Kb if byteSuffix set to true
func ByteToHuman(byteValue int64, byteSuffix bool) string {
	value := int64(0)
	suff := ""

	switch {
	case byteValue >= tb:
		suff = "T"
		value = byteValue / tb
	case byteValue >= gb:
		suff = "G"
		value = byteValue / gb
	case byteValue >= mb:
		suff = "M"
		value = byteValue / mb
	case byteValue >= kb:
		suff = "K"
		value = byteValue / kb
	case byteValue >= b:
		byteSuffix = false
		suff = "B"
		value = byteValue
	}

	strValue := fmt.Sprintf("%s%s", strconv.FormatInt(value, 10), suff)
	if byteSuffix {
		strValue += "b"
	}
	return strValue
}
