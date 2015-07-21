package bytefmt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	b   = float64(1)
	kib = 1024 * b
	mib = 1024 * kib
	gib = 1024 * mib
	tib = 1024 * gib

	kb = 1000 * b
	mb = 1000 * kb
	gb = 1000 * mb
	tb = 1000 * gb
)

// Convert raw string 1B, 1kB, 10MB, 45GiB ... to byte
func ParseString(rawString string) (float64, error) {
	var rawValue string
	var bytes float64
	binaryValue := false

	rawPattern := `^([0-9]+)(\.|\,)?([0-9]+)?([(k|K)MGT])?(i)?(B|b)?$`
	compilePattern := regexp.MustCompile(rawPattern)
	tokens := compilePattern.FindStringSubmatch(rawString)
	// If we passed correct raw string,
	// FindStringSubmatch return slice with 7 elements
	if len(tokens) != 7 {
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

	if tokens[5] == "i" {
		binaryValue = true
	}

	switch strings.ToUpper(tokens[4]) {
	case "B":
		bytes = value * b
	case "K":
		bytes = value * kb
		if binaryValue {
			bytes = value * kib
		}
	case "M":
		bytes = value * mb
		if binaryValue {
			bytes = value * mib
		}
	case "G":
		bytes = value * gb
		if binaryValue {
			bytes = value * gib
		}
	case "T":
		bytes = value * tb
		if binaryValue {
			bytes = value * tib
		}
	default:
		bytes = value * b
	}

	return bytes, nil
}

// Convert bytes to human representation value, like that
// 1024 - 1.02kB if binary is false or 1KiB when it true
func FormatBytes(byteValue float64, prec int, binary bool) string {
	value := float64(0)
	suffix := ""

	if binary {
		switch {
		case byteValue >= tib:
			suffix = "TiB"
			value = byteValue / tib
		case byteValue >= gib:
			suffix = "GiB"
			value = byteValue / gib
		case byteValue >= mib:
			suffix = "MiB"
			value = byteValue / mib
		case byteValue >= kib:
			suffix = "KiB"
			value = byteValue / kib
		case byteValue >= b:
			suffix = "B"
			value = byteValue
		}
	} else {

		switch {
		case byteValue >= tb:
			suffix = "TB"
			value = byteValue / tb
		case byteValue >= gb:
			suffix = "GB"
			value = byteValue / gb
		case byteValue >= mb:
			suffix = "MB"
			value = byteValue / mb
		case byteValue >= kb:
			suffix = "kB"
			value = byteValue / kb
		case byteValue >= b:
			suffix = "B"
			value = byteValue
		}
	}

	strValue := fmt.Sprintf("%s%s",
		strconv.FormatFloat(value, 'f', prec, 64), suffix)
	return strValue
}
