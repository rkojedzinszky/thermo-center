package sensorvalue

import (
	"testing"
)

func TestHTU21D(t *testing.T) {
	var sv SensorValue
	var raw int32

	sv, raw = Parse(63, 255)

	if sv.Metric() != "Temperature" {
		t.Error("Expected Temperature, got:", sv.Metric())
	}

	if raw != 16383 {
		t.Error("Parsed raw mismatch, got:", raw)
	}

	sv, raw = Parse(143, 255)

	if sv.Metric() != "Humidity" {
		t.Error("Expected Humidity, got:", sv.Metric())
	}

	if raw != 4095 {
		t.Error("Parsed raw mismatch, got:", raw)
	}
}

func TestAM302(t *testing.T) {
	var sv SensorValue
	var raw int32

	sv, raw = Parse(159, 255)

	if sv.Metric() != "Temperature" {
		t.Error("Expected Temperature, got:", sv.Metric())
	}

	if raw != -1 {
		t.Error("Parsed raw mismatch, got:", raw)
	}

	sv, raw = Parse(199, 255)

	if sv.Metric() != "Humidity" {
		t.Error("Expected Humidity, got:", sv.Metric())
	}

	if raw != 1023 {
		t.Error("Parsed raw mismatch, got:", raw)
	}
}

func TestOther(t *testing.T) {
	var sv SensorValue
	var raw int32

	sv, raw = Parse(195, 255)

	if sv.Metric() != "Power" {
		t.Error("Expected Power, got:", sv.Metric())
	}

	if raw != 1023 {
		t.Error("Parsed raw mismatch, got:", raw)
	}

	sv, raw = Parse(203, 255)

	if sv.Metric() != "ADC1" {
		t.Error("Expected ADC1, got:", sv.Metric())
	}

	if raw != 1023 {
		t.Error("Parsed raw mismatch, got:", raw)
	}
}

func TestUnknown(t *testing.T) {
	tests := [][]uint8{
		[]byte{255, 255},
		[]byte{64, 0},
		[]byte{160, 0},
	}

	for _, test := range tests {
		rv, _ := Parse(test[0], test[1])

		if rv != nil {
			t.Error("Parsing should've failed")
		}
	}
}
