package sensorvalue

// NewRSSI creates a SensorValue representing RSSI value
func NewRSSI(rssi float32) SensorValue {
	return &plainValue{
		metric: "RSSI",
		value:  float64(rssi),
	}
}

// NewLQI creates a SensorValue representing LQI value
func NewLQI(lqi uint32) SensorValue {
	return &plainValue{
		metric: "LQI",
		value:  float64(lqi),
	}
}

func vccParser(r RawValue) (SensorValue, int32) {
	raw := int32(r.Unsigned)

	return &plainValue{
		metric: "Power",
		value:  1.1 * 1023.0 / float64(raw),
	}, raw
}

func adc1Parser(r RawValue) (SensorValue, int32) {
	raw := int32(r.Unsigned)

	return &plainValue{
		metric: "ADC1",
		value:  float64(raw),
	}, raw
}

func init() {
	register(2, 0, vccParser)
	register(2, 2, adc1Parser)
}
