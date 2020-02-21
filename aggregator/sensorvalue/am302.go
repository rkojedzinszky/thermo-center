package sensorvalue

func am302TemperatureParser(r RawValue) (SensorValue, int32) {
	raw := int32(r.Signed)

	return &plainValue{
		metric: "Temperature",
		value:  float64(raw) / 10.0,
	}, raw
}

func am302HumidityParser(r RawValue) (SensorValue, int32) {
	raw := int32(r.Unsigned)

	return &plainValue{
		metric: "Humidity",
		value:  float64(raw) / 10.0,
	}, raw
}

func init() {
	register(1, 1, am302TemperatureParser)
	register(2, 1, am302HumidityParser)
}
