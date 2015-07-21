package bytefmt

import "testing"

func TestParseString(t *testing.T) {
	var resultValue float64
	resultValue, err := ParseString("4.675TiB")
	if err != nil {
		t.Error(err)
	}
	if int64(resultValue) != 5140216859852 {
		t.Error("Expected 5140216859852. Got ", resultValue)
	}

	resultValue, err = ParseString("2,57GB")
	if err != nil {
		t.Error(err)
	}
	if int64(resultValue) != 2570000000 {
		t.Error("Expected 2570000000. Got ", resultValue)
	}

	resultValue, err = ParseString("2,57")
	if err == nil {
		t.Error("Expected error. Got nil")
	}

}

func TestFormatBytes(t *testing.T) {
	var resultStringValue string

	resultStringValue = FormatBytes(1024, 2, false)
	if resultStringValue != "1.02kB" {
		t.Error("Expected 1.02kB. Got ", resultStringValue)
	}

	resultStringValue = FormatBytes(5140216859852, 3, true)
	if resultStringValue != "4.675TiB" {
		t.Error("Expected 4.675Tb. Got ", resultStringValue)
	}

	resultStringValue = FormatBytes(2570000000, 2, false)
	if resultStringValue != "2.57GB" {
		t.Error("Expected 2.57GB. Got ", resultStringValue)
	}
}
