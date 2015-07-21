package bytefmt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	b  = float64(1)
	kb = 1024 * b
	mb = 1024 * kb
	gb = 1024 * mb
	tb = 1024 * gb
)

// Convert raw string 1B, 1K, 1Kb ... to byte
func ParseString(rawString string) (float64, error) {
	var rawValue string
	var bytes float64
	rawPattern := `^([0-9]+)(\.|\,)?([0-9]+)?([KMGT])?(B|b)?$`
	compilePattern := regexp.MustCompile(rawPattern)
	tokens := compilePattern.FindStringSubmatch(rawString)
	// If we passed correct raw string,
	// FindStringSubmatch return slice with 6 elements
	if len(tokens) != 6 {
		return 0, errors.New("Incorrect string value for conversion")
	}
	switch {
	case tokens[2] == "." || tokens[2] == ",":
		if tokens[4] == "" {
			return 0, errors.New("Incorrect string value for conversion")
		}
		if tokens[2] == "," {
			tokens[2] = "."
		}
		rawValue = strings.Join(tokens[1:4], "")
	default:
		rawValue = tokens[1]
	}
	value, err := strconv.ParseFloat(rawValue, 64)
	if err != nil {
		return value, err
	}

	switch strings.ToUpper(tokens[4]) {
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

// Convert bytes to human representation value, like that
// 1024 - 1K or 1Kb if byteSuffix set to true
func FormatBytes(byteValue float64, prec int, byteSuffix bool) string {
	value := float64(0)
	suffix := ""

	switch {
	case byteValue >= tb:
		suffix = "T"
		value = byteValue / tb
	case byteValue >= gb:
		suffix = "G"
		value = byteValue / gb
	case byteValue >= mb:
		suffix = "M"
		value = byteValue / mb
	case byteValue >= kb:
		suffix = "K"
		value = byteValue / kb
	case byteValue >= b:
		byteSuffix = false
		suffix = "B"
		value = byteValue
	}

	strValue := fmt.Sprintf("%s%s",
		strconv.FormatFloat(value, 'f', prec, 64), suffix)
	if byteSuffix {
		strValue += "b"
	}
	return strValue
}
